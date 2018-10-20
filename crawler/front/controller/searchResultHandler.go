package controller

import (
	"Go_Spider/crawler/front/view"
	"Go_Spider/crawler/model"
	"context"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {

	client, _ := elastic.NewClient(elastic.SetSniff(false))
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

func (handler SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	q := strings.TrimSpace(req.FormValue("q"))
	from, _ := strconv.Atoi(req.FormValue("from"))

	var page model.SeearchResult

	page, _ = handler.GetSearchResult(q, from)

	handler.view.Render(w, page)

}
func (handler SearchResultHandler) GetSearchResult(q string, from int) (model.SeearchResult, error) {

	var result model.SeearchResult

	resp, err := handler.client.Search("dating_profile").Query(elastic.NewQueryStringQuery(q)).From(from).Do(context.Background())
	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Item = resp.Each(reflect.TypeOf(model.Profile{}))

	return result, nil

}
