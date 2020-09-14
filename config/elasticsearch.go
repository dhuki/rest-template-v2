package config

import (
	"fmt"

	"github.com/dhuki/rest-template-v2/common"
	"github.com/elastic/go-elasticsearch/v7"
)

// add this "discovery.type: single-node" in elasticsearch.yml config to avoid the production check
// while run elaticsearch server

func NewElasticsearch() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("http://%s:%s", common.ElasticsHost, common.ElasticsPort),
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return client, nil
}
