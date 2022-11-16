package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Elton-Bezerra/ports-and-adapter/adapters/dto"
	"github.com/Elton-Bezerra/ports-and-adapter/app"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service app.ProductServiceInterface) {
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")
}

func getProduct(service app.ProductServiceInterface) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func createProduct(service app.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var productDto dto.Product

		err := json.NewDecoder(r.Body).Decode(&productDto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonError(err.Error()))
			return
		}

		product, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}
