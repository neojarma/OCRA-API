package elasticsearch_repository

import (
	"bytes"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type ElasticsearchRepository interface {
	Find(req *bytes.Buffer) (map[string]interface{}, error)
	AutoComplete(req *bytes.Buffer) (map[string]interface{}, error)
	AddDocument(req *esapi.IndexRequest) error
}
