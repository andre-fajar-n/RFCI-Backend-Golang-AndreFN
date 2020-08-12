package main

import (
	"fmt"
	"html"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Message struct {
	Message string `json:message`
}

type Data struct {
	Author string `json:"author"`
	Title string `json:"title"`
	Comments []Message `json:"comments"`
}
var data = Data{}

func PostMessage(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&data)

	json.NewEncoder(w).Encode(&data)
}

func main()  {
	router := mux.NewRouter()

	router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	router.HandleFunc("/message", PostMessage).Methods("POST")

	http.ListenAndServe(":8000", router)
}