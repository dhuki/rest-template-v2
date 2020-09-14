package usecase

import (
	"context"

	"github.com/dhuki/rest-template-v2/common"
	"github.com/dhuki/rest-template-v2/pkg/testing/domain/entity"
	"github.com/dhuki/rest-template-v2/pkg/testing/domain/repo"
	"github.com/dhuki/rest-template-v2/utils"
)

// in usecaseImpl struct we just inject functionality of dependecy (dependent to abstraction (DI (Dependency Inversion)))
// to make it plug and play
// not directly dependency itself (not dependent to lower module)
type UsecaseImpl struct {
	TestTableRepo repo.TestTableRepo
	Utils         utils.Utils
}

// mat ryer style just pass few argument as paramater to any kind func
// restrict it to max 3 argument
// func NewUsecase(testTableRepo repo.TestTableRepo, utils utils.Utils) Usecase {
// 	return UsecaseImpl{
// 		testTableRepo: testTableRepo,
// 		utils:         utils,
// 	}
// }

func (u UsecaseImpl) GetAllData(ctx context.Context) common.BaseResponse {
	var response common.BaseResponse
	{
		// time.Sleep(3 * time.Second)
		testTables, err := u.TestTableRepo.GetAll(ctx)
		if err != nil {
			return common.BaseResponse{
				Error: err,
			}
		}

		response.Data = testTables
	}
	response.Success = common.RESPONSE_SUCCESS
	response.Message = common.RESPONSE_MSG_SUCCESS

	return response
}

func (u UsecaseImpl) CreateData(ctx context.Context, request entity.TestTable) common.BaseResponse {
	err := u.TestTableRepo.Create(ctx, request)
	if err != nil {
		return common.BaseResponse{
			Error: err,
		}
	}

	// try to send email
	go func() {
		u.Utils.SendEmail(ctx, request)
	}()

	return common.BaseResponse{
		Success: common.RESPONSE_SUCCESS,
		Message: common.RESPONSE_MSG_SUCCESS,
		Data:    nil,
		Error:   nil,
	}
}

func (u UsecaseImpl) GetDataByParam(ctx context.Context, request entity.TestTable) common.BaseResponse {
	testTable, err := u.TestTableRepo.Get(ctx, request.ID)
	if err != nil {
		return common.BaseResponse{
			Error: err,
		}
	}

	return common.BaseResponse{
		Success: common.RESPONSE_SUCCESS,
		Message: common.RESPONSE_MSG_SUCCESS,
		Data:    testTable,
		Error:   nil,
	}
}

func (u UsecaseImpl) GetDataByPath(ctx context.Context, request entity.TestTable) common.BaseResponse {
	testTable, err := u.TestTableRepo.Get(ctx, request.ID)
	if err != nil {
		return common.BaseResponse{
			Error: err,
		}
	}

	return common.BaseResponse{
		Success: common.RESPONSE_SUCCESS,
		Message: common.RESPONSE_MSG_SUCCESS,
		Data:    testTable,
		Error:   nil,
	}
}
