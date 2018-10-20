package persist

import (
	"Go_Spider/crawler/model"
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() chan interface{} {

	itemChan := make(chan interface{})

	client, _ := elastic.NewClient(
		elastic.SetSniff(false))

	go func() {
		itemCount := 0
		for true {
			item := <-itemChan
			log.Printf("ItemSaver:    Got item #%d  %v ", itemCount, item)
			itemCount++

			switch item.(type) {
			case model.Profile:
				profile := item.(model.Profile)
				_, err := save(client, profile.Id, profile)
				if err != nil {
					log.Printf("Save Item Error :    saving item: %v    error:   %v", profile, err)
				}
			}

		}
	}()

	return itemChan

}

func save(client *elastic.Client, id string, item interface{}) (reId string, err error) {

	response, err := client.Index().Index("dating_profile").Type("zhenai").Id(id).BodyJson(item).Do(context.Background())

	if err != nil {
		return "", err
	}

	fmt.Printf("%+v", response)

	return response.Id, nil

}

func get(id string) {

	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	result, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	var actual model.Profile
	err = json.Unmarshal(*result.Source, &actual)

	if err != nil {
		panic(err)
	}

}
