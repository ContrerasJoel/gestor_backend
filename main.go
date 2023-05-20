package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ContrerasJoel/gestor_go/db"
	"github.com/ContrerasJoel/gestor_go/internal/product"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	DSN := os.Getenv("DSN")
	if DSN == "" {
		DSN = "postgres://root:8sWlwpSrlR93tUUlrShNEhE8PPQAnBfj@dpg-chj5l9l269v2e2eeoglg-a.oregon-postgres.render.com/gestor_dvgi"
	}
	db.ConnectionDB(DSN)
	db.DB.AutoMigrate(product.Product{})

	port := os.Getenv("PORT")
	if port == "" {
		port = "7071"
	}

	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requests-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	product.NewHandler().Expose(r)

	srv := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: handlers.CORS(headers, methods, origins)(r),
	}

	log.Printf("Escuchando en el puerto %s", srv.Addr)
	srv.ListenAndServe()
}
