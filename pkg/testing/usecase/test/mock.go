package test

import (
	"context"

	"github.com/dhuki/rest-template-v2/pkg/testing/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockTestTableRepo struct {
	mock.Mock
}

func (m MockTestTableRepo) GetAll(ctx context.Context) ([]entity.TestTable, error) {
	// Called tells the mock object that a method has been called, and gets an array
	// of arguments to return
	args := m.Called(ctx)
	return args.Get(0).([]entity.TestTable), args.Error(1)
}

func (m MockTestTableRepo) Get(ctx context.Context, id int) (entity.TestTable, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.TestTable), args.Error(1)
}

func (m MockTestTableRepo) GetByName(ctx context.Context, name string) (entity.TestTable, error) {
	return entity.TestTable{}, nil
}

func (m MockTestTableRepo) Create(ctx context.Context, data entity.TestTable) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

type MockEmail struct {
	mock.Mock
}

func (m MockEmail) SendEmail(ctx context.Context, data entity.TestTable) error {
	return nil
}
