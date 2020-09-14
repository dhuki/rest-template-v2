package test

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dhuki/rest-template-v2/config"
	"github.com/dhuki/rest-template-v2/pkg/testing/domain/entity"
	"github.com/dhuki/rest-template-v2/pkg/testing/infrastructure/command"
	"github.com/stretchr/testify/assert"
)

// this is testing with mock connection database by sqlmock lib
// this sqlmock lib only check if query is correct with actual query or not
// it will pass if :
// 1. expected query contain actual query, regardless queries exactly same or not with actual query
// 2. number of dynamic argument is same
// 3. return query muxt be same

// command to get coverprofile if test in different package
// go test -coverpkg ./... ./test -coverprofile ./test/coverage.out
// to show in html format
// go tool cover -html=cp.out

func TestUnitGetAll(t *testing.T) {
	// using this default mathcer so it will using regex
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	// expected query
	mock.ExpectQuery(`[SELECT (.+) FROM test_tables]`).
		// will output from query above
		WillReturnRows(sqlmock.NewRows([]string{"Name", "Description"}).
			AddRow("testing", "testing"))

	gormDB, err := config.NewTesting(db)
	assert.Nil(t, err)

	repo := command.NewTestTableCommand(gormDB)
	_, err = repo.GetAll(context.TODO())
	assert.Nil(t, err)

	// we make sure that all expectations were met
	err = mock.ExpectationsWereMet()
	assert.Nil(t, err)
}

func TestUnitGet(t *testing.T) {
	// using this QueryMatcherEqual to ignore regex and become native query for postgres in this case
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.Nil(t, err)
	defer db.Close()

	// expected query
	mock.ExpectQuery(`SELECT * FROM "test_tables" WHERE "test_tables"."id" = $1 LIMIT 1`).
		WithArgs(1).
		// will output from query above
		WillReturnRows(sqlmock.NewRows([]string{"Name", "Description"}).
			AddRow("testing", "testing"))

	gormDB, err := config.NewTesting(db)
	assert.Nil(t, err)

	repo := command.NewTestTableCommand(gormDB)
	_, err = repo.Get(context.TODO(), 1)
	assert.Nil(t, err)

	// we make sure that all expectations were met
	err = mock.ExpectationsWereMet()
	assert.Nil(t, err)
}

func TestUnitGetByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	expected := entity.TestTable{
		ID:          1,
		Name:        "testing",
		Description: "testing",
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "test_tables" WHERE name = $1`)).
		WithArgs("testing").
		WillReturnRows(
			sqlmock.NewRows([]string{"ID", "Name", "Description"}).
				AddRow(1, "testing", "testing"))

	gormDB, err := config.NewTesting(db)
	assert.Nil(t, err)

	repo := command.NewTestTableCommand(gormDB)
	actual, err := repo.GetByName(context.TODO(), "testing")
	assert.Nil(t, err)

	// we make sure that all expectations were met
	assert.Nil(t, mock.ExpectationsWereMet())

	assert.Equal(t, actual, expected, "it's suppose to be same")
}

func TestUnitCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	// all these query is works
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "test_tables" ("name","description") VALUES ($1,$2) RETURNING "id"`)).
		WithArgs("testing", "testing").
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	// mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "test_tables" ("name","description") VALUES ($1,$2) RETURNING`)).
	// 	WithArgs("testing", "testing").
	// 	WillReturnRows(sqlmock.NewRows([]string{"id"}))

	// mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "test_tables" ("name","description")`)).
	// 	WithArgs("testing", "testing").
	// 	WillReturnRows(sqlmock.NewRows([]string{"id"}))

	gormDB, err := config.NewTesting(db)
	assert.Nil(t, err)

	repo := command.NewTestTableCommand(gormDB)
	err = repo.Create(context.TODO(), entity.TestTable{
		Name:        "testing",
		Description: "testing",
	})
	assert.Nil(t, err)

	// we make sure that all expectations were met
	err = mock.ExpectationsWereMet()
	assert.Nil(t, err)
}
