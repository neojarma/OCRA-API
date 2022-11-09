package elasticsearch_service

import (
	"bytes"
	"encoding/json"
	model "ocra_server/model/entity"
	elasticsearch_repository "ocra_server/repository/elasticsearch"
	"os"
	"sync"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type ElasticsearchServiceImpl struct {
	Repo elasticsearch_repository.ElasticsearchRepository
}

func NewElasticsearchService(repo elasticsearch_repository.ElasticsearchRepository) ElasticsearchService {
	var doOnce sync.Once
	service := new(ElasticsearchServiceImpl)

	doOnce.Do(func() {
		service = &ElasticsearchServiceImpl{
			Repo: repo,
		}
	})

	return service
}

func (Service *ElasticsearchServiceImpl) Find(query string) ([]string, error) {
	var buf bytes.Buffer
	body := map[string]interface{}{
		"_source": []string{"videoId"},
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":     query,
				"fuzziness": "auto",
				"fields":    []string{"videoTitle", "videoDesc", "videoTags"},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, err
	}

	res, err := Service.Repo.Find(&buf)
	if err != nil {
		return nil, err
	}

	return Service.parsingBodyResult(res, "videoId"), nil
}

func (Service *ElasticsearchServiceImpl) AutoComplete(query string) ([]string, error) {
	var buf bytes.Buffer
	body := map[string]interface{}{
		"_source": []string{"videoTitle"},
		"query": map[string]interface{}{
			"prefix": map[string]interface{}{
				"videoTitle": query,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, err
	}

	res, err := Service.Repo.Find(&buf)
	if err != nil {
		return nil, err
	}

	return Service.parsingBodyResult(res, "videoTitle"), nil

}

func (Service *ElasticsearchServiceImpl) AddDocument(request *model.ElasticsearchVideo) error {
	bodyByte, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index: os.Getenv("ES_INDEX"),
		Body:  bytes.NewReader(bodyByte),
	}

	return Service.Repo.AddDocument(&req)
}

func (Service *ElasticsearchServiceImpl) parsingBodyResult(res map[string]interface{}, field string) []string {
	var id []string
	for _, hit := range res["hits"].(map[string]interface{})["hits"].([]interface{}) {
		each := hit.(map[string]interface{})["_source"].(map[string]interface{})[field].(string)
		id = append(id, each)
	}

	return id
}
