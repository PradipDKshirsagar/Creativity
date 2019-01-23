package interest

import (
	"fmt"
	"database/sql"
    "creativity/db"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)


type Interest struct {
    ID int `json:"id"`
    Name string `json:"name"`
}

//type Interests []Interest


func Create(w http.ResponseWriter, r *http.Request) {
    db := db.Db 
    
    sqlStatement := `INSERT INTO interests (name) VALUES ($1) RETURNING id`
    
    var data Interest   
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        fmt.Println("err....")
    }
    fmt.Println(data)
    id := 0
    
    err = db.QueryRow(sqlStatement, data.Name).Scan(&id)
    if err != nil {
        panic(err)
    }
    fmt.Println("New record Interest ID is:", id)
    fmt.Fprint(w,"Your user ID is:", id, "\nDo not forgot")
}

func Read(w http.ResponseWriter, r *http.Request) {
    db := db.Db 
    sqlStatement := `SELECT * FROM interests WHERE ID = $1`

    vars := mux.Vars(r)
    id := vars["id"]

    var in Interest
    row := db.QueryRow(sqlStatement, id)

    err := row.Scan(&in.ID,&in.Name)
    switch err {
    case sql.ErrNoRows:
        fmt.Println("interest id not exist")
        fmt.Fprint(w,"interest id not exist")
        return
    case nil:
        fmt.Println(in)
        b, err := json.Marshal(in)
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
    sqlStatement := `UPDATE interests SET name = $2 WHERE id = $1;`

    vars := mux.Vars(r)
    id := vars["id"]

    var data Interest    
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        fmt.Println("err....")
    }
    fmt.Println(data)

    res, err := db.Exec(sqlStatement, id, data.Name)
    if err != nil {
      panic(err)
    }
    cnt, err := res.RowsAffected()
    if err != nil {
      panic(err)
    }
    fmt.Println(cnt)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    db := db.Db 
    sqlStatement := `DELETE FROM interests WHERE id = $1;`
    
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
    sqlStatement := `SELECT * FROM interests;`

    var interests []Interest

    rows, err := db.Query(sqlStatement)
    if err != nil {
        fmt.Fprint(w,"No interests list  present")    
    }
    defer rows.Close()
    var in Interest
    for rows.Next() {
        err = rows.Scan(&in.ID, &in.Name)
        if err != nil {
            panic(err)
        }
        fmt.Println(in)
        interests = append(interests,in)
    }
    b, err := json.Marshal(interests)
    if err != nil {
        fmt.Printf("Error: %s", err)
        return;
    }
    fmt.Fprint(w,string(b))
}   

func Check(id int) bool {
    db := db.Db
    sqlStatement := `select * from interests where id = $1;`

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