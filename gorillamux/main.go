package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	rootRouter := mux.NewRouter()

	apiV1Router := rootRouter.PathPrefix("/api/v1").Subrouter()
	apiV1Router.HandleFunc("/", func(resWriter http.ResponseWriter, req *http.Request) {
		fmt.Fprint(resWriter, "gorilla mux")
	}).Methods(http.MethodGet)
}
