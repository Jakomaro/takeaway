package repository

import "github.com/elastic/go-elasticsearch/v8"

func NewElasticDB() (*elasticsearch.Client, error) {

	return elasticsearch.NewDefaultClient()
}
