package interest

import (
	"fmt"
	"database/sql"
)

type Interest struct {
    ID int 
    Name string
}


func Create(db *sql.DB) {
    var name string
    
    sqlStatement := `INSERT INTO interests (name) VALUES ($1) RETURNING id`
    
    fmt.Println("Ente Interest Name : ",)
    fmt.Scan(&name)
    
    id := 0
    
    err := db.QueryRow(sqlStatement, name).Scan(&id)
    if err != nil {
        panic(err)
    }
    fmt.Println("New record Interest ID is:", id)
}

func Read(db *sql.DB) {
    sqlStatement := `SELECT * FROM interests WHERE ID = $1`

    var uid int
    fmt.Scan(&uid)
    var in Interest
    
    row := db.QueryRow(sqlStatement, uid)

    err := row.Scan(&in.ID, &in.Name)
    switch err {
    case sql.ErrNoRows:
      fmt.Println("No rows were returned!")
      return
    case nil:
      fmt.Println(in)
    default:
      panic(err)
    }
}

func Update(db *sql.DB) {
    sqlStatement := `
    UPDATE interests SET name = $2 WHERE id = $1;`

    var name string
    var id int

    fmt.Println("Enter User ID which want to update details : ")
    fmt.Scan(&id)
    fmt.Println("Enter the Interest (Name) : ")
    fmt.Scan(&name)
    res, err := db.Exec(sqlStatement, id, name)
    if err != nil {
      panic(err)
    }
    _, err = res.RowsAffected()
    if err != nil {
      panic(err)
    }
    //fmt.Println(count)
}

func Delete(db *sql.DB) {
    sqlStatement := `DELETE FROM interests WHERE id = $1;`
    
    var id int

    fmt.Println("Enter the interest id which want to delete : ")
    fmt.Scan(&id)
    _, err := db.Exec(sqlStatement, id)
    if err != nil {
      panic(err)
    }
}