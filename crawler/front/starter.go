package main

import (
	"Go_Spider/crawler/front/controller"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("crawler/front/view")))
	http.Handle("/search", controller.CreateSearchResultHandler("crawler/front/view/search.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
