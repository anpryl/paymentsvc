package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/anpryl/paymentsvc/services"
	"github.com/anpryl/paymentsvc/svcerrors"
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
)

func New(as services.AccountService) http.Handler {
	r := chi.NewRouter()
	listAccountsHandler := httptransport.NewServer(
		accounts(as),
		decodeListAccountsReq,
		encodeResponse,
	)
	r.Method(http.MethodPost, "/accounts", listAccountsHandler)
	return r
}

func decodeListAccountsReq(_ context.Context, r *http.Request) (interface{}, error) {
	strLimit := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(strLimit)
	if err != nil {
		return nil, svcerrors.ErrInvalidLimitValue
	}
	strOffset := r.URL.Query().Get("offset")
	offset, err := strconv.Atoi(strOffset)
	if err != nil {
		return nil, svcerrors.ErrInvalidOffsetValue
	}
	return ListAccountsReq{
		Limit:  limit,
		Offset: offset,
	}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type ListAccountsReq struct {
	Limit  int
	Offset int
}

func accounts(svc services.AccountService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(ListAccountsReq)
		return svc.ListOfAccounts(req.Offset, req.Limit)
	}
}
