package product

import (
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (s *Handler) Expose(r *mux.Router) {
	r.HandleFunc("/api/product/{id}", s.ReadProduct).Methods("GET")
	r.HandleFunc("/api/products", s.ReadProducts).Methods("GET")
	r.HandleFunc("/api/product", s.CreateProduct).Methods("POST")
	r.HandleFunc("/api/product/{id}", s.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/product/{id}", s.DeleteProduct).Methods("DELETE")
}
