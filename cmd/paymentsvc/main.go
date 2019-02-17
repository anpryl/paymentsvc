package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anpryl/paymentsvc/api"
	"github.com/anpryl/paymentsvc/config"
	"github.com/anpryl/paymentsvc/services"
)

func main() {
	cfg := config.FromEnv()
	as := services.NewAccount()
	r := api.New(as)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Http.Host, cfg.Http.Port), r)
	log.Fatal(err)
}
