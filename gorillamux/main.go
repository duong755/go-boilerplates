package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	rootRouter := mux.NewRouter()

	apiV1Router := rootRouter.PathPrefix("/api/v1").Subrouter()

	// swagger:route GET / demo demoRoot
	//
	// Return a string
	//
	//
	//     Consumes:
	//     - application/json
	//     - application/text
	//     - application/xml
	//
	//     Produces:
	//     - application/text
	//
	//     Schemes: http, https
	//
	//     Deprecation: false
	//
	//
	apiV1Router.HandleFunc("/", func(resWriter http.ResponseWriter, req *http.Request) {
		fmt.Fprint(resWriter, "gorilla mux")
	}).Methods(http.MethodGet)

	http.ListenAndServe(":3333", rootRouter)
}
