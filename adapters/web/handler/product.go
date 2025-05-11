package handler

import (
	"arquitetura-hexagonal-go/application"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeProductHandlers(r mux.Router, n negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(request)
		id := vars["id"]
		fmt.Printf("\nbuscando pelo produto %s\n", id)
		product, err := service.Get(id)
		if err != nil {
			response.WriteHeader(http.StatusNotFound)
			log.Fatalf("error to proccess request. %s", err.Error())
			return
		}

		err = json.NewEncoder(response).Encode(product)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			log.Fatalf("erro to convert product to json: %s", err.Error())
			return
		}
	})
}