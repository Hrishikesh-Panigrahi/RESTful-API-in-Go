package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}
var Articles []Article

func AllArticles(w http.ResponseWriter, r *http.Request){
	// articles := Articles{
	// 	Article{Title:"Test title", Desc:"Test description", Content:"Hello world"},
	// }
	fmt.Println("Article endpoint hit")
	fmt.Fprintln(w, "hello")
	json.NewEncoder(w).Encode(Articles)
}

func PostArticles(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "post Endpoint hit")
}
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home page Endpoint hit")
}

func handleRequest(){
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/articles", AllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", PostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", myRouter))
}

func main(){

	Articles = []Article{
        Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }

	handleRequest()
}