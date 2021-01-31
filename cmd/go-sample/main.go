package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-sample/pkg/articles"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"Hello World!"}`))
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", hello)
	r.HandleFunc("/articles", articles.CreateArticle).Methods("POST")
	r.HandleFunc("/articles", articles.ReturnArticles)
	r.HandleFunc("/articles/{id}", articles.UpdateArticle).Methods("PUT")
	r.HandleFunc("/articles/{id}", articles.DeleteArticle).Methods("DELETE")
	r.HandleFunc("/articles/{id}", articles.ReturnArticle)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	fmt.Println("Rest API v1.8 - Mux Routers")
	handleRequests()
}
