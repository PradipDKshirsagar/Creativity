package user

import (
	"fmt"
    "creativity/db"
    "encoding/json"
    "net/http"
    "database/sql"
    "github.com/gorilla/mux"
)

type User struct {
    ID int `json: "id"`
    Age int `json:"age"`
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
}

type Users []User

func Create(w http.ResponseWriter, r *http.Request) {
    db := db.Db 
    
    sqlStatement := `INSERT INTO users (age, first_name, last_name) VALUES ($1, $2, $3) RETURNING id`
    var data User 
    
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        fmt.Println("err....")
    }
    fmt.Println(data)
    
    id := 0
    
    err = db.QueryRow(sqlStatement, data.Age, data.FirstName, data.LastName).Scan(&id)
    if err != nil {
        panic(err)
    }
    fmt.Println("New record ID is:", id)
    fmt.Fprint(w,"Your interest ID is:", id, "\nDo not forgot")
}


func Read(w http.ResponseWriter, r *http.Request) {
    db := db.Db 
    sqlStatement := `SELECT * FROM USERS WHERE ID = $1;`

    vars := mux.Vars(r)
    id := vars["id"]

    var ur User
    row := db.QueryRow(sqlStatement, id)

    err := row.Scan(&ur.ID, &ur.Age, &ur.FirstName,&ur.LastName)
    switch err {
    case sql.ErrNoRows:
        fmt.Println("user id not exist")
        fmt.Fprint(w,"user id not exist")
        return
    case nil:
        fmt.Println(ur)
        b, err := json.Marshal(ur)
        if err != nil {
            fmt.Printf("Error: %s", err)
            return;
        }
        fmt.Fprint(w,string(b))
    default:
        panic(err)
    }
}

func Update(w http.ResponseWriter, r *http.Request) {
    db := db.Db 
    sqlStatement := `UPDATE users SET age = $2, first_name = $3, last_name = $4 WHERE id = $1;`
    fmt.Println("working")

    vars := mux.Vars(r)
    id := vars["id"]

    var data User   
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        fmt.Println("err....")
    }
    fmt.Println(data)

    res, err := db.Exec(sqlStatement, id,data.Age, data.FirstName, data.LastName)
    if err != nil {
      panic(err)
    }
    count, err := res.RowsAffected()
    if err != nil {
      panic(err)
    }
    fmt.Println(count)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    db := db.Db 
    sqlStatement := `DELETE FROM users WHERE id = $1;`
    
    vars := mux.Vars(r)
    id := vars["id"]

    res, err := db.Exec(sqlStatement, id)
    if err != nil {
      panic(err)
    }
    cnt, err := res.RowsAffected()
    if err != nil {
      panic(err)
    }
    fmt.Println(cnt)
}

func ReadAll(w http.ResponseWriter, r *http.Request) {
    db := db.Db 
    sqlStatement := `SELECT * FROM USERS;`

    var users Users

    rows, err := db.Query(sqlStatement)
    if err != nil {
        fmt.Fprint(w,"Err to retrive users")
        return    
    }

    defer rows.Close()
    var ur User
    for rows.Next() {
        err = rows.Scan(&ur.ID, &ur.Age, &ur.FirstName,&ur.LastName)
        if err != nil {
            panic(err)
        }
        fmt.Println(ur)
        users = append(users,ur)
    }
    b, err := json.Marshal(users)
    if err != nil {
        fmt.Printf("Error: %s", err)
        return;
    }
    fmt.Fprint(w,string(b))
}


func Check(id int) bool {
    db := db.Db
    sqlStatement := `select * from users where id = $1;`

    res, err := db.Exec(sqlStatement, id)
    if err != nil {
      panic(err)
    }
    count, err := res.RowsAffected()
    if err != nil {
      panic(err)
    }
    fmt.Println(count)

    if count == 0 {
       
        return false
    }
    return true 
}