package presenter

import (
	"net/http"

	"github.com/dhuki/rest-template-v2/common"
	"github.com/dhuki/rest-template-v2/middleware"
	"github.com/dhuki/rest-template-v2/pkg/testing/presenter/reqres"
	"github.com/dhuki/rest-template-v2/pkg/testing/usecase"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func NewHttpHandler(r *httprouter.Router, usecase usecase.Usecase, middlewares []alice.Constructor, logger log.Logger) {
	// alice package make chaining middleware more easy
	// before -> middlware1(middleware2(actualFunc))
	// after -> alice.New(middlewares1, middleware2).Then(actualFunc)
	commonMiddleware := alice.New(middlewares...)

	options := []httptransport.ServerOption{
		httptransport.ServerBefore(httptransport.PopulateRequestContext),      // executed on the HTTP request object before the request is decoded.
		httptransport.ServerErrorEncoder(middleware.ErrorEncoder),             // error for client.
		httptransport.ServerErrorHandler(middleware.ErrorHandlerFunc(logger)), // error for internal.
		httptransport.ServerFinalizer(middleware.SetInterceptors(logger)),     // executed at the end of every HTTP request.
	}

	// order is not matter in httprouter
	// bcs it will find more precise/exact than gorilla mux router or default server mux router
	// and more faster than gorilla mux

	// httprouter can work with http.handler built in package
	r.Handler(http.MethodGet, common.BaseUrl+"/testing/:param", commonMiddleware.Then(httptransport.NewServer(
		MakeGetDataByPathEndpointWithGoroutine(usecase),
		reqres.DecodeGetByPathRequest,
		httptransport.EncodeJSONResponse,
		options...,
	)))

	// httprouter cannot do this
	// r.Methods(http.MethodGet).Path("/testing").Queries(
	// 	"param", "{param:[0-9]+}", // [0-9]+ match all number
	// ).Handler(httptransport.NewServer(
	// 	MakeGetDataByParamEndpointWithGoroutine(usecase),
	// 	reqres.DecodeGetByParamRequest,
	// 	httptransport.EncodeJSONResponse,
	// 	options...,
	// ))

	r.Handler(http.MethodGet, common.BaseUrl+"/testing", commonMiddleware.Then(httptransport.NewServer(
		MakeGetAllDataEndpointWithGoroutine(usecase),
		httptransport.NopRequestDecoder,  // go-kit provided requests that do not need to be decoded
		httptransport.EncodeJSONResponse, // go-kit provided response to be encoded to json and add req header as content-type json
		options...,
	)))

	// r.Methods(http.MethodPost).Path("/testing").Handler(httptransport.NewServer(
	// 	MakeCreateDataEndpoint(usecase),
	// 	middleware.LogRequestBody(reqres.DecodeCreateRequest, logger), // request need to be decoded
	// 	httptransport.EncodeJSONResponse,
	// 	options...,
	// ))
}
