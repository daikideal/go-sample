package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
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

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"Hello World!"}`))
}

func handleRequests() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/articles", returnArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Server started.")
	createDemoArticles()
	handleRequests()
}