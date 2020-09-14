package usecase

import (
	"context"

	"github.com/dhuki/rest-template-v2/common"
	"github.com/dhuki/rest-template-v2/pkg/testing/domain/entity"
)

type Usecase interface {
	GetAllData(context.Context) common.BaseResponse
	GetDataByParam(context.Context, entity.TestTable) common.BaseResponse
	GetDataByPath(context.Context, entity.TestTable) common.BaseResponse
	CreateData(context.Context, entity.TestTable) common.BaseResponse
}
