package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Handler) ReadProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	productOne, err := ReadOne(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{"error": "Product No Found"})
	} else {
		json.NewEncoder(w).Encode(productOne)
	}
}

func (s *Handler) ReadProducts(w http.ResponseWriter, r *http.Request) {
	productsAll, rowsAffected := ReadAll()
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]any{"error": "Products No Found"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productsAll)
}

func (s *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	json.NewDecoder(r.Body).Decode(&product)
	createdProduct := CreateOne(&product)
	if createdProduct != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{"error": createdProduct.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&product)

}

func (s *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	product := Product{}
	json.NewDecoder(r.Body).Decode(&product)
	updatedProduct, err := UpdateOne(product, params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{"error": "Product No Found"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedProduct)
}

func (s *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	deletedProduct := DeleteOne(params["id"])
	fmt.Println(deletedProduct)
	if deletedProduct == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]any{"error": "Product No Found"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{"status": "delete success"})

}
