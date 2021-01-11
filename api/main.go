package main

import (
	"fmt"
	"log"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Article struct {
	Id int `json:"id"`
	Title	string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func createDemoArticles() {
	Articles = []Article{
		Article{Id: 1, Title: "Test Article", Desc: "Article Description", Content: "Test Content"},
		Article{Id: 2, Title: "Test Article 2", Desc: "Article Description 2", Content: "Test Content 2"},
	}
}

func returnArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}

func returnArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Article Id", http.StatusBadRequest)
	}

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	var article Article

	json.NewDecoder(r.Body).Decode(&article)
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Article Id", http.StatusBadRequest)
	}

	var updateArticle Article
	json.NewDecoder(r.Body).Decode(&updateArticle)

	for index, article := range Articles {
		if article.Id == key {
			Articles[index] = updateArticle
			json.NewEncoder(w).Encode(updateArticle)
		}
	}
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Article Id", http.StatusBadRequest)
	}

	for index, article := range Articles {
		if article.Id == key {
			Articles = append(Articles[:index], Articles[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"Hello World!"}`))
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", hello)
	r.HandleFunc("/articles", createArticle).Methods("POST")
	r.HandleFunc("/articles", returnArticles)
	r.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	r.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	r.HandleFunc("/articles/{id}", returnArticle)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	fmt.Println("Rest API v1.8 - Mux Routers")
	createDemoArticles()
	handleRequests()
}