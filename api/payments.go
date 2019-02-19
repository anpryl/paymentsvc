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
	uuid "github.com/satori/go.uuid"
)

func paymentEndpoints(r *chi.Mux, ps services.Payments) {
	accountPaymentsHandler := httptransport.NewServer(
		accountPaymentsEndpoint(ps),
		decodeAccountPaymentsRequest,
		encodeResponseOK,
	)
	r.Method(http.MethodGet, "/accounts/{accountID}/payments", accountPaymentsHandler)
	createPaymentHander := httptransport.NewServer(
		createPaymentEndpoint(ps),
		decodeCreatePaymentReq,
		encodeResponseCreated,
	)
	r.Method(http.MethodPost, "/payments", createPaymentHander)
}

func createPaymentEndpoint(svc services.Payments) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(services.NewPayment)
		id, err := svc.CreatePayment(ctx, req)
		return IDResp{ID: id}, err
	}
}

func decodeCreatePaymentReq(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var p services.NewPayment
	err := json.NewDecoder(r.Body).Decode(&p)
	return p, err
}

type AccountPaymentsReq struct {
	models.OffsetLimit
	ID uuid.UUID
}

func accountPaymentsEndpoint(svc services.Payments) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AccountPaymentsReq)
		return svc.AccountPayments(ctx, req.ID, req.OffsetLimit)
	}
}

func decodeAccountPaymentsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req AccountPaymentsReq
	var err error
	req.OffsetLimit, err = decodeOffsetLimitRequest(r)
	if err != nil {
		return nil, err
	}
	id, err := uuid.FromString(chi.URLParam(r, "accountID"))
	if err != nil {
		return nil, err
	}
	req.ID = id
	return req, nil
}
