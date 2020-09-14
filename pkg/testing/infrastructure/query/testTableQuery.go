package query

import "github.com/elastic/go-elasticsearch/v7"

type TestTableQuery interface {
	Search() error
}

type TestTableQueryImpl struct {
	Elasticlient *elasticsearch.Client
}

func NewTestTableQuery(elasticlient *elasticsearch.Client) TestTableQuery {
	return TestTableQueryImpl{
		Elasticlient: elasticlient,
	}
}

func (t TestTableQueryImpl) Search() error {
	return nil
}
