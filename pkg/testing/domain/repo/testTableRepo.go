package repo

import (
	"github.com/dhuki/rest-template-v2/pkg/testing/infrastructure/command"
	"github.com/dhuki/rest-template-v2/pkg/testing/infrastructure/query"
)

// package for defined repository to work with database (try to implement architectural pattern CQRS)
// try to implement interface segregation (“Clients should not be forced to depend upon interfaces that they do not use.”)
// source : https://code.tutsplus.com/tutorials/solid-part-3-liskov-substitution-interface-segregation-principles--net-36710

// type TestTableRepo interface {
// 	command.TestTableCommand
// 	query.TestTableQuery
// }

// using this implementation bcs the interface has DI (depency injection) inside its function.
type TestTableRepo struct {
	command.TestTableCommand
	query.TestTableQuery
}

func NewTestTableRepo(command command.TestTableCommand, query query.TestTableQuery) TestTableRepo {
	return TestTableRepo{
		TestTableCommand: command,
		TestTableQuery:   query,
	}
}
