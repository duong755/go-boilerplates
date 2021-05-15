package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

// swagger:response successResponse
type successResponse = string

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
	//     Response:
	//       200: successResponse
	//
	apiV1Router.HandleFunc("/", func(resWriter http.ResponseWriter, req *http.Request) {
		resWriter.Header().Add("Content-Type", "application/json")
		resWriter.WriteHeader(http.StatusOK)
		resWriter.Write([]byte("gorilla mux"))
	}).Methods(http.MethodGet)

	specDocument, specErr := loads.Spec("./swagger.json")
	if specErr != nil {
		log.Fatal(specErr)
	}
	swaggerJSON, marshalErr := json.MarshalIndent(specDocument.Pristine().Spec(), "", "  ")
	if marshalErr != nil {
		log.Fatal(marshalErr)
	}

	apiV1Router.HandleFunc("/swagger.json", func(resWriter http.ResponseWriter, req *http.Request) {
		resWriter.Header().Add("Content-Type", "application/json")
		resWriter.WriteHeader(http.StatusOK)
		resWriter.Write(swaggerJSON)
	})
	swaggerOpts := middleware.SwaggerUIOpts{SpecURL: "/api/v1/swagger.json", BasePath: "/api/v1", Path: "docs"}
	swaggerMiddleware := middleware.SwaggerUI(swaggerOpts, nil)
	apiV1Router.Handle("/docs", swaggerMiddleware)

	http.ListenAndServe(":3333", rootRouter)
}
