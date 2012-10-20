package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/product/{key}", ProductHandler)
	r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	r.HandleFunc("/id={id:[0-9]+}", ArticleIdHandler)
	http.Handle("/", r)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("Home request")
}

func ProductHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("Product request")
	vars := mux.Vars(req)
	key := vars["key"]
	if key != "" {
		log.Println("key =", key)
	}
}

func ArticlesCategoryHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("Category request")
	vars := mux.Vars(req)
	category := vars["category"]
	if category != "" {
		log.Println("category =", category)
	}
}

func ArticleHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("Article request")
	vars := mux.Vars(req)
	category := vars["category"]
	if category != "" {
		log.Println("category =", category)
		id := vars["id"]
		if id != "" {
			log.Println("id =", id)
		}
	}
}

func ArticleIdHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("Article ID request")
	vars := mux.Vars(req)
	id := vars["id"]
	if id != "" {
		log.Println("id =", id)
	}
}
