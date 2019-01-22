package main 

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func main() {
	r :=mux.NewRouter()
	r.HandleFunc("/ping", pingHandler)
	http.Handle("/",r)
	fmt.Println("Starting Server at port 8080")
	http.ListenAndServe(":8080", nil)	
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"PONG")
}