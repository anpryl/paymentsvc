package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/anpryl/paymentsvc/models"
	"github.com/anpryl/paymentsvc/services"
	"github.com/anpryl/paymentsvc/svcerrors"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	uuid "github.com/satori/go.uuid"
)

func New(as services.AccountService) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	accountsEndpoints(r, as)
	return r
}

type IDResp struct {
	ID uuid.UUID `json:"id"`
}

func decodeOffsetLimitReq(_ context.Context, r *http.Request) (interface{}, error) {
	var err error
	var req models.OffsetLimit
	strLimit := r.URL.Query().Get("limit")
	if strLimit == "" {
		req.Limit = models.DefaultLimit
	} else {
		req.Limit, err = strconv.Atoi(strLimit)
		if err != nil {
			return nil, svcerrors.ErrInvalidLimitValue
		}
	}
	strOffset := r.URL.Query().Get("offset")
	if strOffset == "" {
		req.Offset = models.DefaultOffset
	} else {
		req.Offset, err = strconv.Atoi(strOffset)
		if err != nil {
			return nil, svcerrors.ErrInvalidOffsetValue
		}
	}
	return req, nil
}

func encodeResponseOK(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

func encodeResponseCreated(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(response)
}
