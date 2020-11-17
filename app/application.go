package app

import (
	"net/http"
	"time"

	"github.com/angadthandi/bookstore_items-api/clients/elasticsearch"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()

	mapUrls()

	serverConfig := &http.Server{
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  2 * time.Second,
		Handler:      router,
	}

	if err := serverConfig.ListenAndServe(); err != nil {
		panic(err)
	}
}
