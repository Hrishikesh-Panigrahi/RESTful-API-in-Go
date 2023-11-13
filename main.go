package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}
type Articles []Article

func AllArticles(w http.ResponseWriter, r *http.Request){
	articles := Articles{
		Article{Title:"Test title", Desc:"Test description", Content:"Hello world"},
	}
	fmt.Println("Article endpoint hit")
	json.NewEncoder(w).Encode(articles)
}


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home page Endpoint hit")
}

func handleRequest(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", AllArticles)
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func main(){
	handleRequest()
}