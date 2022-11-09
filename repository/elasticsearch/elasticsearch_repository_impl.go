package elasticsearch_repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"os"
	"sync"

	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var INDEX = os.Getenv("ES_INDEX")

type ElasticsearchRepositoryImpl struct {
	ES *elasticsearch.Client
}

func NewElasticsearchRepository(es *elasticsearch.Client) ElasticsearchRepository {
	var doOnce sync.Once
	repo := new(ElasticsearchRepositoryImpl)

	doOnce.Do(func() {
		repo = &ElasticsearchRepositoryImpl{
			ES: es,
		}
	})

	return repo
}

func (repo *ElasticsearchRepositoryImpl) Find(req *bytes.Buffer) (map[string]interface{}, error) {
	res, err := repo.ES.Search(
		repo.ES.Search.WithIndex(INDEX),
		repo.ES.Search.WithBody(req),
		repo.ES.Search.WithTrackTotalHits(true),
	)

	if err != nil {
		return nil, err
	}

	return repo.parsingBodyResult(res)
}

func (repo *ElasticsearchRepositoryImpl) AutoComplete(req *bytes.Buffer) (map[string]interface{}, error) {
	res, err := repo.ES.Search(
		repo.ES.Search.WithIndex(INDEX),
		repo.ES.Search.WithBody(req),
		repo.ES.Search.WithTrackTotalHits(true),
	)

	if err != nil {
		return nil, err
	}

	return repo.parsingBodyResult(res)
}

func (repo *ElasticsearchRepositoryImpl) AddDocument(req *esapi.IndexRequest) error {
	res, err := req.Do(context.Background(), repo.ES)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return errors.New("error" + res.Status())
	}

	return nil
}

func (repo *ElasticsearchRepositoryImpl) parsingBodyResult(res *esapi.Response) (map[string]interface{}, error) {
	if res.IsError() {
		var e map[string]interface{}

		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, errors.New(fmt.Sprintln("error parsing", err))
		}

		return nil, fmt.Errorf("[%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"],
		)
	}

	var r map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, errors.New(fmt.Sprintln("error parsing", err))
	}

	return r, nil
}
