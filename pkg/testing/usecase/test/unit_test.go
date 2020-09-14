package test

import (
	"context"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/dhuki/rest-template-v2/common"
	"github.com/dhuki/rest-template-v2/pkg/testing/domain/entity"
	"github.com/dhuki/rest-template-v2/pkg/testing/domain/repo"
	"github.com/dhuki/rest-template-v2/pkg/testing/usecase"
	"github.com/dhuki/rest-template-v2/utils"
	"github.com/stretchr/testify/assert"
)

// command to get coverprofile if test in different package
// go test -coverpkg ./... ./test -coverprofile ./test/coverage.out
// to show in html format
// go tool cover -html=cp.out

func TestGetAllData(t *testing.T) {
	mockCommand := new(MockTestTableRepo)
	mockRepo := repo.NewTestTableRepo(mockCommand, nil)

	mockEmail := new(MockEmail)
	utils := utils.NewUtils().WireWithEmail(mockEmail)

	var data entity.TestTable
	assert.Nil(t, faker.FakeData(&data))

	mockCommand.On("GetAll", context.TODO()).Return([]entity.TestTable{data}, nil)

	usecase := usecase.UsecaseImpl{
		TestTableRepo: mockRepo,
		Utils:         utils,
	}
	actual := usecase.GetAllData(context.TODO())
	assert.Nil(t, actual.Error)

	expected := common.BaseResponse{
		Success: true,
		Message: "Success",
		Data:    []entity.TestTable{data},
		Error:   nil,
	}

	assert.Equal(t, expected, actual, "it's supposed to be same")
}

func TestGetDataByParam(t *testing.T) {
	mockCommand := new(MockTestTableRepo)
	mockRepo := repo.NewTestTableRepo(mockCommand, nil)

	mockEmail := new(MockEmail)
	mockUtils := utils.NewUtils().WireWithEmail(mockEmail)

	var expected entity.TestTable
	assert.Nil(t, faker.FakeData(&expected))

	mockCommand.On("Get", context.TODO(), expected.ID).Return(expected, nil)

	usecase := usecase.UsecaseImpl{
		TestTableRepo: mockRepo,
		Utils:         mockUtils,
	}

	actual := usecase.GetDataByParam(context.TODO(), expected)
	assert.Nil(t, actual.Error)

	assert.Equal(t, expected, actual.Data, "it's supposed to be same")
}

func TestGetDataByPath(t *testing.T) {
	mockCommand := new(MockTestTableRepo)
	mockRepo := repo.NewTestTableRepo(mockCommand, nil)

	mockEmail := new(MockEmail)
	mockUtils := utils.NewUtils().WireWithEmail(mockEmail)

	var expected entity.TestTable
	assert.Nil(t, faker.FakeData(&expected))

	mockCommand.On("Get", context.TODO(), expected.ID).Return(expected, nil)

	usecase := usecase.UsecaseImpl{
		TestTableRepo: mockRepo,
		Utils:         mockUtils,
	}

	actual := usecase.GetDataByPath(context.TODO(), expected)
	assert.Nil(t, actual.Error)

	assert.Equal(t, expected, actual.Data, "it's supposed to be same")
}

func TestCreateData(t *testing.T) {
	mockCommand := new(MockTestTableRepo)
	mockRepo := repo.NewTestTableRepo(mockCommand, nil)

	mockEmail := new(MockEmail)
	mockUtils := utils.NewUtils().WireWithEmail(mockEmail)

	var request entity.TestTable
	assert.Nil(t, faker.FakeData(&request))

	mockCommand.On("Create", context.TODO(), request).Return(nil)

	usecase := usecase.UsecaseImpl{
		TestTableRepo: mockRepo,
		Utils:         mockUtils,
	}

	actual := usecase.CreateData(context.TODO(), request)
	assert.Nil(t, actual.Error)
}
