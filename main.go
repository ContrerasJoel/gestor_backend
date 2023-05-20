package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/ContrerasJoel/gestor_go/db"
	"github.com/ContrerasJoel/gestor_go/internal/product"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DSN := os.Getenv("DSN")
	db.ConnectionDB(DSN)
	db.DB.AutoMigrate(product.Product{})

	url := "https://api.ipify.org?format=json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(string(content))

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
