package api

import (
	"context"
	"net/http"

	"github.com/anpryl/paymentsvc/services"
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
)

func currenciesEndpoints(r *chi.Mux, cs services.Currency) {
	listAccountsHandler := httptransport.NewServer(
		allCurrenciesEndpoint(cs),
		httptransport.NopRequestDecoder,
		encodeResponseOK,
	)
	r.Method(http.MethodGet, "/currencies", listAccountsHandler)
}

func allCurrenciesEndpoint(svc services.Currency) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.AllCurrencies(ctx)
	}
}
