package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()

	serverConfig := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
	}

	if err := serverConfig.ListenAndServe(); err != nil {
		panic(err)
	}
}
