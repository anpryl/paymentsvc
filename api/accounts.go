package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/anpryl/paymentsvc/models"
	"github.com/anpryl/paymentsvc/services"
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
)

func accountsEndpoints(r *chi.Mux, as services.AccountService) {
	listAccountsHandler := httptransport.NewServer(
		accountsEndpoint(as),
		decodeOffsetLimitReq,
		encodeResponseOK,
	)
	r.Method(http.MethodGet, "/accounts", listAccountsHandler)
	addAccountHandler := httptransport.NewServer(
		addAccountEndpoint(as),
		decodeAddAccountReq,
		encodeResponseCreated,
	)
	r.Method(http.MethodPost, "/accounts", addAccountHandler)
}

func addAccountEndpoint(svc services.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(services.NewAccount)
		id, err := svc.AddAccount(ctx, req)
		return IDResp{ID: id}, err
	}
}

func decodeAddAccountReq(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var acc services.NewAccount
	err := json.NewDecoder(r.Body).Decode(&acc)
	return acc, err
}

func accountsEndpoint(svc services.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.OffsetLimit)
		return svc.ListOfAccounts(ctx, req)
	}
}
