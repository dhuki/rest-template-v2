package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dhuki/rest-template-v2/common"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
)

// function for interceptors
func SetInterceptors(logger log.Logger) httptransport.ServerFinalizerFunc {
	return httptransport.ServerFinalizerFunc(func(ctx context.Context, code int, r *http.Request) {
		level.Info(logger).Log(
			"description", "Interceptors",
			"scheme", ctx.Value(httptransport.ContextKeyRequestProto),
			"host", ctx.Value(httptransport.ContextKeyRequestHost),
			"path", ctx.Value(httptransport.ContextKeyRequestPath),
			"method", ctx.Value(httptransport.ContextKeyRequestMethod),
			"statusCode", code)
		// "requestBody", fmt.Sprintf("%+v", request)) // givin output of struct to this -> attribute : value
		// "response", fmt.Sprintf("%+v", w.)) // givin output of struct to this -> attribute : value
	})
}

// function for logging error to client
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var response common.BaseResponse
	{
		switch err {
		case common.ErrDataNotFound:
			response.Data = common.ErrDataNotFound.Error()
		case common.ErrAssertion:
			response.Data = common.ErrAssertion.Error()
		case common.ErrCancelled:
			w.WriteHeader(http.StatusOK)
			response.Data = common.ErrCancelled.Error()
		case common.ErrLimitExceed:
			response.Data = common.ErrLimitExceed.Error()
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	json.NewEncoder(w).Encode(response)
}

// function for logging error to internal
func ErrorHandlerFunc(logger log.Logger) transport.ErrorHandlerFunc {
	return transport.ErrorHandlerFunc(func(_ context.Context, err error) {
		level.Error(logger).Log(
			"description", "Internal server error",
			"message", err,
			"solution", "Please check encode/decode body, usecase method, and dependency library",
		)
	})
}
