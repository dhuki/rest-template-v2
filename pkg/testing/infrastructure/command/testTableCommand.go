package command

import (
	"context"

	"github.com/dhuki/rest-template-v2/pkg/testing/domain/entity"
	"gorm.io/gorm"
)

// package that implement repo in domain layer
// it places outermost (paling luar) of all layer

type TestTableCommand interface {
	GetAll(context.Context) ([]entity.TestTable, error)
	Get(context.Context, int) (entity.TestTable, error)
	GetByName(context.Context, string) (entity.TestTable, error)
	Create(context.Context, entity.TestTable) error
}

type TestTableCommandImpl struct {
	db *gorm.DB
}

func NewTestTableCommand(db *gorm.DB) TestTableCommand {
	return TestTableCommandImpl{
		db: db,
	}
}

func (t TestTableCommandImpl) GetAll(ctx context.Context) ([]entity.TestTable, error) {
	var testTables []entity.TestTable
	// db := t.db.Find(&testTables) // this is ver1 of gorm cannot use context
	db := t.db.WithContext(ctx).Find(&testTables) // this is ver2 of gorm, we can use context to provide cancellation propagation
	if db.Error != nil {
		return nil, db.Error
	}
	return testTables, nil
}

func (t TestTableCommandImpl) Get(ctx context.Context, id int) (entity.TestTable, error) {
	var testTables entity.TestTable
	db := t.db.WithContext(ctx).Take(&testTables, id) // only work with int of type primary key
	if db.Error != nil {
		return testTables, db.Error
	}

	return testTables, nil
}

func (t TestTableCommandImpl) GetByName(ctx context.Context, name string) (entity.TestTable, error) {
	var testTables entity.TestTable
	// db := t.db.Find(&testTables) // this is ver1 of gorm cannot use context
	db := t.db.WithContext(ctx).Where("name = ?", name).First(&testTables) // this is ver2 of gorm, we can use context to provide cancellation propagation
	if db.Error != nil {
		return testTables, db.Error
	}

	return testTables, nil
}

func (t TestTableCommandImpl) Create(ctx context.Context, testTable entity.TestTable) error {
	db := t.db.WithContext(ctx).Create(&testTable) // this is ver2 of gorm, we can use context to provide cancellation propagation
	if db.Error != nil {
		return db.Error
	}
	return nil
}
