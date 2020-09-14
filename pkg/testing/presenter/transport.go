package presenter

import (
	"net/http"

	"github.com/dhuki/rest-template-v2/middleware"
	"github.com/dhuki/rest-template-v2/pkg/testing/presenter/reqres"
	"github.com/dhuki/rest-template-v2/pkg/testing/usecase"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpHandler(mux *mux.Router, usecase usecase.Usecase, middlewares []mux.MiddlewareFunc, logger log.Logger) {
	r := mux.PathPrefix("/api").Subrouter()
	// r.Use(middleware.SetContentTypeToJson) // by default go-kit provided function to response as json
	// r.Use(middleware.SetInterceptors(logger)) // by default there go-kit provided "ServerFinalizerFunc" for logging request
	// r.Use(handlers.CompressHandler) // compress response to gzip encoding
	r.Use(middlewares...)

	options := []httptransport.ServerOption{
		httptransport.ServerBefore(httptransport.PopulateRequestContext),      // executed on the HTTP request object before the request is decoded.
		httptransport.ServerErrorEncoder(middleware.ErrorEncoder),             // error for client.
		httptransport.ServerErrorHandler(middleware.ErrorHandlerFunc(logger)), // error for internal.
		httptransport.ServerFinalizer(middleware.SetInterceptors(logger)),     // executed at the end of every HTTP request.
	}

	// order matter in mux match router

	r.Methods(http.MethodGet).Path("/testing/{param}").Handler(httptransport.NewServer(
		MakeGetDataByPathEndpointWithGoroutine(usecase),
		reqres.DecodeGetByPathRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods(http.MethodGet).Path("/testing").Queries(
		"param", "{param:[0-9]+}", // [0-9]+ match all number
	).Handler(httptransport.NewServer(
		MakeGetDataByParamEndpointWithGoroutine(usecase),
		reqres.DecodeGetByParamRequest,
		httptransport.EncodeJSONResponse,
		options...,
	))

	r.Methods(http.MethodGet).Path("/testing").Handler(httptransport.NewServer(
		MakeGetAllDataEndpointWithGoroutine(usecase),
		httptransport.NopRequestDecoder,  // go-kit provided requests that do not need to be decoded
		httptransport.EncodeJSONResponse, // go-kit provided response to be encoded to json and add req header as content-type json
		options...,
	))

	r.Methods(http.MethodPost).Path("/testing").Handler(httptransport.NewServer(
		MakeCreateDataEndpoint(usecase),
		middleware.LogRequestBody(reqres.DecodeCreateRequest, logger), // request need to be decoded
		httptransport.EncodeJSONResponse,
		options...,
	))
}
