package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/anpryl/paymentsvc/migrations"
	"github.com/anpryl/paymentsvc/repositories"

	"github.com/anpryl/paymentsvc/api"
	"github.com/anpryl/paymentsvc/config"
	"github.com/anpryl/paymentsvc/services"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
)

func main() {
	cfg := config.FromEnv()
	db := connectDB(cfg)
	txFunc := repositories.NewWithTransactionFunc(db)
	tryMigrate(db)
	ar := repositories.NewAccountsRepository(db)
	er := repositories.NewExchangeRatesRepository(db)
	cr := repositories.NewCurrenciesRepository(db)
	pr := repositories.NewPaymentsRepository(db)
	as := services.NewAccountsService(ar, cr)
	cs := services.NewCurrenciesService(cr)
	ps := services.NewPaymentsService(pr, ar, cr, er, txFunc)
	r := api.New(as, cs, ps)
	log.Println("Starting server")
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Http.Host, cfg.Http.Port), r)
	log.Fatal(err)
}

func connectDB(cfg config.Config) *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.PostgreSQL.Host, cfg.PostgreSQL.Port),
		User:     cfg.PostgreSQL.User,
		Password: cfg.PostgreSQL.Password,
		Database: cfg.PostgreSQL.Database,
	})
}

func tryMigrate(db *pg.DB) {
	_, err := db.Exec(`
      	CREATE TABLE IF NOT EXISTS gopg_migrations (
      	       id serial,
      	       version bigint,
      	       created_at timestamptz
	)`)
	if err != nil {
		log.Fatalf("Failed to create gopg_migrations table: %v", err)
	}
	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}
