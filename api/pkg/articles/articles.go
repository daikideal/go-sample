package articles

import (
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

func init() {
	createDemoArticles()
}

func createDemoArticles() {
	Articles = []Article{
		Article{Id: 1, Title: "Test Article", Desc: "Article Description", Content: "Test Content"},
		Article{Id: 2, Title: "Test Article 2", Desc: "Article Description 2", Content: "Test Content 2"},
	}
}

func ReturnArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Articles)
}

func ReturnArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Article Id", http.StatusBadRequest)
	}

	for _, article := range Articles {
		if article.Id == key {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(article)
		}
	}
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article Article

	json.NewDecoder(r.Body).Decode(&article)
	Articles = append(Articles, article)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
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
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateArticle)
		}
	}
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Article Id", http.StatusBadRequest)
	}

	for index, article := range Articles {
		if article.Id == key {
			Articles = append(Articles[:index], Articles[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		}
	}
}