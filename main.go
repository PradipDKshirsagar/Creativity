package main 

import (
	"fmt"
    "net/http"
    "github.com/gorilla/mux"
	 "creativity/db"
	"creativity/interest"
    "creativity/usersinterests"
    "creativity/user"    
)
func pingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w,"PONG")
}

func main() {
    db.DBConnetion()
    defer db.Db.Close()

    r := mux.NewRouter()

    r.HandleFunc("/ping", pingHandler)

    r.HandleFunc("/user", user.Create).Methods("POST")
    r.HandleFunc("/user/{id}", user.Delete).Methods("DELETE")
    r.HandleFunc("/user/{id}", user.Read).Methods("GET")
    r.HandleFunc("/user/{id}", user.Update).Methods("PUT")
    r.HandleFunc("/user", user.ReadAll).Methods("GET")

    r.HandleFunc("/interest", interest.Create).Methods("POST")
    r.HandleFunc("/interest/{id}", interest.Delete).Methods("DELETE")
    r.HandleFunc("/interest/{id}", interest.Read).Methods("GET")
    r.HandleFunc("/interest/{id}", interest.Update).Methods("PUT")
    r.HandleFunc("/interest", interest.ReadAll).Methods("GET")

    r.HandleFunc("/usersinterests", usersinterests.Add).Methods("POST")
    r.HandleFunc("/usersinterests/{uid}",usersinterests.Check).Methods("GET")

    http.Handle("/",r)
    http.ListenAndServe(":8080", nil)
}	