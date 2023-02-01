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
type SuccessResponse struct {
	// in: body
	Body string
}

func main() {
	rootRouter := mux.NewRouter()

	fileServer := http.FileServer(http.Dir("./static"))
	rootRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	v1Router(rootRouter)
	swaggerRouter(rootRouter)

	println("Listening on port 3333 - https://localhost:3333/docs")
	certFile := "./localhost.pem"
	keyFile := "./localhost-key.pem"
	http.ListenAndServeTLS(":3333", certFile, keyFile, rootRouter)
}

func v1Router(rootRouter *mux.Router) {
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
	//     Responses:
	//       200: successResponse
	//
	apiV1Router.HandleFunc("/", func(resWriter http.ResponseWriter, req *http.Request) {
		resWriter.Header().Add("Content-Type", "application/text")
		resWriter.WriteHeader(http.StatusOK)
		resWriter.Write([]byte("gorilla mux"))
	}).Methods(http.MethodGet)
}

func swaggerRouter(rootRouter *mux.Router) {
	specDocument, specErr := loads.Spec("./swagger.json")
	if specErr != nil {
		log.Fatal(specErr)
	}
	swaggerJSON, marshalErr := json.MarshalIndent(specDocument.Pristine().Spec(), "", "  ")
	if marshalErr != nil {
		log.Fatal(marshalErr)
	}

	rootRouter.HandleFunc("/swagger.json", func(resWriter http.ResponseWriter, req *http.Request) {
		resWriter.Header().Add("Content-Type", "application/json")
		resWriter.WriteHeader(http.StatusOK)
		resWriter.Write(swaggerJSON)
	})
	swaggerOpts := middleware.SwaggerUIOpts{
		SpecURL:  "/swagger.json",
		BasePath: "/",
		Path:     "docs",
		Title:    "Demo API Documentation",
	}
	swaggerMiddleware := middleware.SwaggerUI(swaggerOpts, nil)
	rootRouter.Handle("/docs", swaggerMiddleware)
}
