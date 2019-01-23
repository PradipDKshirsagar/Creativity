package interest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Interest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	var data Interest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("err....")
		return
	}
	fmt.Println(data)

	err = createService(data)

	if err != nil {
		fmt.Fprint(w, "Some error is occured %v", err)
		return
	}
	fmt.Fprint(w, "Successfully insert..")
}

func Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	interest, err := readService(id)
	if err != nil {
		fmt.Fprint(w, "Error is ", err)
		return
	}
	fmt.Println(interest)
	b, err := json.Marshal(interest)
	if err != nil {
		fmt.Fprint(w, "not marshaling")
		return
	}
	fmt.Fprint(w, string(b))
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var data Interest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("err....")
	}
	fmt.Println(data)

	err = updateService(id, data)

	if err != nil {
		fmt.Fprint(w, "Error is ", err)
		return
	}
	fmt.Fprint(w, "Successfully updated...")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := deleteService(id)

	if err != nil {
		fmt.Fprint(w, "Error is ", err)
		return
	}
	fmt.Fprint(w, "Successfully Delete user...")
}

func ReadAll(w http.ResponseWriter, r *http.Request) {

	interests, err := readAllService()
	if err != nil {
		fmt.Fprint(w, "Error is ", err)
	}
	b, err := json.Marshal(interests)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(interests)
	fmt.Fprint(w, string(b))
}
