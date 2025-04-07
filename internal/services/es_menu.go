package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/jakomaro/takeaway/internal/model"
)

type ESMenuService struct {
	es *elasticsearch.Client
}

func NewESMenuService(es *elasticsearch.Client) *ESMenuService {
	return &ESMenuService{es: es}
}

func (es *ESMenuService) GetMenu(ctx context.Context) ([]model.Item, error) {

	query := `{ "query": {"match_all": {}} }`
	res, err := es.es.Search(
		es.es.Search.WithContext(context.Background()),
		es.es.Search.WithIndex("menu"),
		es.es.Search.WithBody(bytes.NewReader([]byte(query))),
		es.es.Search.WithTrackTotalHits(true),
		es.es.Search.WithPretty(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error in response: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search request failed with status: %d", res.StatusCode)
	}

	var result map[string]any
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response body: %v", err)
	}

	total := result["hits"].(map[string]any)["total"].(map[string]any)
	tot := total["value"].(float64)
	menu := make([]model.Item, int(tot))

	hits := result["hits"].(map[string]any)["hits"].([]any)
	for i, hit := range hits {
		doc := hit.(map[string]any)["_source"]
		// fmt.Printf("%v\n", doc)
		row := doc.(map[string]any)
		name := row["name"].(string)
		price := row["price"].(float64)
		item := model.Item{
			ItemID: i + 1,
			Name:   name,
			Price:  price,
		}
		menu[i] = item
	}

	return menu, nil
}
