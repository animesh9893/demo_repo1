package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"os"
)

type Post struct {
  Message string `json:"Message"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Hello Animesh\n");
}

func HomeGetHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin","*")
	var p Post
	p.Message = "Hello Unknown"
	json.NewEncoder(w).Encode(p)
}

func HomePostHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin","*")
	var p Post
	decoder := json.NewDecoder(r.Body)
	err:= decoder.Decode(&p)
	if err!=nil{
		p.Message = "Something wrong"
	}
	p.Message = "All Good"

	json.NewEncoder(w).Encode(p)
}

func main(){
	port := os.Getenv("PORT")
	r:= mux.NewRouter()
	r.HandleFunc("/",HomeHandler)
	r.HandleFunc("/home",HomeGetHandler).Methods("GET")
	r.HandleFunc("/home",HomePostHandler).Methods("POST")
	http.ListenAndServe(":"+port,r)
}