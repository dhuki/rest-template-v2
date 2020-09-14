package presenter

import (
	"context"

	"github.com/dhuki/rest-template-v2/common"
	"github.com/dhuki/rest-template-v2/pkg/testing/domain/entity"
	"github.com/dhuki/rest-template-v2/pkg/testing/usecase"
	"github.com/go-kit/kit/endpoint"
)

func MakeGetAllDataEndpointWithGoroutine(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// source : https://www.ardanlabs.com/blog/2018/11/goroutine-leaks-the-forgotten-sender.html
		// if we're using unbuffered channel it will lead to leak
		response := make(chan common.BaseResponse, 1)

		go func() {
			response <- usecase.GetAllData(ctx)
		}()

		select {
		case <-ctx.Done():
			return nil, common.ErrCancelled
		case result := <-response:
			if result.Error != nil {
				return nil, result.Error
			}
			return result, nil
		}
	}
}

func MakeGetAllDataEndpoint(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		response := usecase.GetAllData(ctx)
		return response, response.Error
	}
}

func MakeCreateDataEndpoint(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(entity.TestTable)
		if !ok {
			return nil, common.ErrAssertion
		}
		response := usecase.CreateData(ctx, req)
		return response, response.Error
	}
}

func MakeGetDataByParamEndpointWithGoroutine(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		testTable, ok := request.(entity.TestTable)
		if !ok {
			return nil, common.ErrAssertion
		}

		responseChan := make(chan common.BaseResponse, 1)

		go func() {
			responseChan <- usecase.GetDataByParam(ctx, testTable)
		}()

		select {
		case <-ctx.Done():
			return nil, common.ErrCancelled
		case response := <-responseChan:
			if response.Error != nil {
				return nil, response.Error
			}
			return response, nil
		}
	}
}

func MakeGetDataByPathEndpointWithGoroutine(usecase usecase.Usecase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		testTable, ok := request.(entity.TestTable)
		if !ok {
			return nil, common.ErrAssertion
		}

		responseChan := make(chan common.BaseResponse, 1)

		go func() {
			responseChan <- usecase.GetDataByPath(ctx, testTable)
		}()

		select {
		case <-ctx.Done():
			return nil, common.ErrCancelled
		case response := <-responseChan:
			if response.Error != nil {
				return nil, response.Error
			}
			return response, nil
		}
	}
}
