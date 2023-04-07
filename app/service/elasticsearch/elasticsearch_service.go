package elasticsearch_service

import (
	model "ocra_server/model/entity"
)

type ElasticsearchService interface {
	Find(query string) ([]string, error)
	AutoComplete(query string) ([]string, error)
	AddDocument(request *model.ElasticsearchVideo) error
}
