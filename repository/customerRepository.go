package repository

import (
	"context"
	db2 "elastic_search_customers/db"
	"elastic_search_customers/models"
	"elastic_search_customers/utils"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
)

func GetCustomerBetweenDates(clientId string, startDate string, endDate string) ([]models.DataStruct, error) {
	ctx := context.Background()
	esclient, err := db2.GetESClient()
	if err != nil {
		return nil, err
	}

	var dataStructs []models.DataStruct

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("ClientID", clientId))
	searchSource.Query(elastic.NewRangeQuery("Time").Gte(startDate).Lte(endDate))

	indexes := utils.BuildIndex("customer_", strings.Split(startDate, "T")[0], strings.Split(endDate, "T")[0])
	fmt.Println(indexes)
	searchService := esclient.Search().Index(indexes).SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		return nil, err
	}

	for _, hit := range searchResult.Hits.Hits {
		var dataStruct models.DataStruct
		err := json.Unmarshal(hit.Source, &dataStruct)
		if err != nil {
			fmt.Println("[Getting DataStruct][Unmarshal] Err=", err)
		}

		dataStructs = append(dataStructs, dataStruct)
	}
	return dataStructs, err
}
