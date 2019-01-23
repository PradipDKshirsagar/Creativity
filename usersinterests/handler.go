package usersinterests

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type usersInterests struct {
	UsID int `json:"user_id"`
	InID int `json:"interest_id"`
}

func Check(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["uid"]

	usIn, err := checkService(id)

	b, err := json.Marshal(usIn)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Fprint(w, string(b))
}

func Add(w http.ResponseWriter, r *http.Request) {
	var data usersInterests
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("err....")
	}
	fmt.Println(data)
	err = addService(data)
	if err != nil {
		fmt.Fprint(w, "Error is ", err)
		return
	}
	fmt.Fprint(w, "Interest added successfully.")
}
