package user

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
)

type User struct {
    ID int `json: "id"`
    Age int `json:"age"`
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	var data User 
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        fmt.Println("err....")
        return 
    }
    fmt.Println(data)

    err = createService(data)

    if err!=nil {
    	fmt.Fprint(w,"Some error is occured %v", err)
    	return
    }
    fmt.Fprint(w, "Successfully insert..")
  //  fmt.Println("New record ID is:", id)
    //fmt.Fprint(w,"Your interest ID is:", id, "\nDo not forgot")
}

func Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id := vars["id"]
    user, err := readService(id)
    if err != nil {
    	fmt.Fprint(w,"Error is ",err)
    	return
    }
    fmt.Println(user)
    b, err := json.Marshal(user)
    if err != nil {           
        fmt.Fprint(w,"not marshaling")
        return 
    }
    fmt.Fprint(w, string(b))
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id := vars["id"]
   	var data User
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        fmt.Println("err....")
    }
    fmt.Println(data)

    err = updateService(id,data)

    if err != nil {
    	fmt.Fprint(w,"Error is ",err)
    	return
    }
    fmt.Fprint(w,"Successfully updated...")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id := vars["id"]

    err := deleteService(id)

    if err != nil {
    	fmt.Fprint(w,"Error is ",err)
    	return
    }
    fmt.Fprint(w,"Successfully Delete user...")
}

func ReadAll(w http.ResponseWriter, r *http.Request) {
	
	//var Users []User
	users, err := readAllService()
	if err !=nil {
		fmt.Fprint(w,"Error is ",err)
	}
	b, err := json.Marshal(users)
    if err != nil {
        fmt.Printf("Error: %s", err)
        return;
    }
    fmt.Println(users)
    fmt.Fprint(w,string(b))
}