package usersinterests

import (
	"fmt"
	"creativity/db"
	"creativity/interest"
    "creativity/user"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

type usersInterests struct {
	UsID int `json:"user_id"`
	InID int `json:"interest_id"`
}

func Check(w http.ResponseWriter, r *http.Request) {
	db := db.Db 
    sqlStatement := `SELECT * FROM interests where Id in (select interest_id from users_interests where user_id = $1);`

    vars := mux.Vars(r)
    id := vars["uid"]

    var usIn interest.Interests
    rows, err := db.Query(sqlStatement,id)
    if err != nil {
        fmt.Fprint(w,"No interests present")    
    }
    defer rows.Close()
    var in interest.Interest
    for rows.Next() {
        err = rows.Scan(&in.ID, &in.Name)
        if err != nil {
            panic(err)
        }
        fmt.Println(in)
        usIn = append(usIn,in)
    }
    b, err := json.Marshal(usIn)
    if err != nil {
        fmt.Printf("Error: %s", err)
        return;
    }
    fmt.Fprint(w,string(b))
}

func Add(w http.ResponseWriter, r *http.Request) {
	db := db.Db

	var data usersInterests   
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        fmt.Println("err....")
    }
    fmt.Println(data)

	if !user.Check(data.UsID) {
		fmt.Fprint(w,"Invalid user id")
		return 
	} 
	if !interest.Check(data.InID){
		fmt.Fprint(w,"Invalid interest id")
		return
	}

	sqlStatement := `INSERT INTO users_interests (user_id, interest_id) VALUES ($1, $2) RETURNING id`

	id := 0
    err = db.QueryRow(sqlStatement, data.UsID, data.InID).Scan(&id)
    if err != nil {
        panic(err)
    }

	fmt.Fprint(w,"Interest added successfully.")
}