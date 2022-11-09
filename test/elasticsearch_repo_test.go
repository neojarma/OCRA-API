package test

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"strings"
	"testing"

	esrepo "ocra_server/repository/elasticsearch"

	"github.com/elastic/go-elasticsearch/v8"
)

func TestElasticsearchRepo(t *testing.T) {
	config := elasticsearch.Config{
		Addresses: []string{"http://localhost:9201", "http://localhost:9202", "http://localhost:9203"},
	}

	es, err := elasticsearch.NewClient(config)
	if err != nil {
		log.Println(err)
	}

	esRepository := esrepo.NewElasticsearchRepository(es)

	// body request
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":     "dota 2",
				"fuzziness": "auto",
				"fields":    []string{"videoTitle", "videoDesc", "videoTags"},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Println(err)
	}

	res, err := esRepository.Find(&buf)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(res)
	}

}

func TestParsing(t *testing.T) {
	env := os.Getenv("ES_ADDRESSES")
	log.Println(env)

	res := strings.Split(env, ",")

	log.Println(res)
}
