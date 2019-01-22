package user

import (
	"fmt"
	"database/sql"
)

type user struct {
    ID int 
    Age int
    FirstName string
    LastName string
}

func Create(db *sql.DB) {
    var firstName, lastName string
    var age int
    
    sqlStatement := `INSERT INTO users (age, first_name, last_name) VALUES ($1, $2, $3) RETURNING id`
    
    fmt.Println("Enter First Name : ",)
    fmt.Scan(&firstName)
    fmt.Println("Enter Last Name : ",)
    fmt.Scan(&lastName)
    fmt.Println("Enter Age : ",)
    fmt.Scan(&age)
    
    id := 0
    
    err := db.QueryRow(sqlStatement, age, firstName, lastName).Scan(&id)
    if err != nil {
        panic(err)
    }
    fmt.Println("New record ID is:", id)
}

func Read(db *sql.DB) {
    sqlStatement := `SELECT * FROM USERS WHERE ID = $1`

    var uid int
    fmt.Scan(&uid)
    var ur user
    
    row := db.QueryRow(sqlStatement, uid)

    err := row.Scan(&ur.ID, &ur.Age, &ur.FirstName,&ur.LastName)
    switch err {
    case sql.ErrNoRows:
      fmt.Println("No rows were returned!")
      return
    case nil:
      fmt.Println(ur)
    default:
      panic(err)
    }
}

func Update(db *sql.DB) {
    sqlStatement := `
    UPDATE users
    SET first_name = $2, last_name = $3
    WHERE id = $1;`

    var firstName, lastName string
    var id int

    fmt.Println("Enter User ID which want to update details : ")
    fmt.Scan(&id)
    fmt.Println("Enter the updated First Name : ")
    fmt.Scan(&firstName)
    fmt.Println("Enter the updated Last Name : ")
    fmt.Scan(&lastName)
    res, err := db.Exec(sqlStatement, id,firstName, lastName)
    if err != nil {
      panic(err)
    }
    count, err := res.RowsAffected()
    if err != nil {
      panic(err)
    }
    fmt.Println(count)
}

func Delete(db *sql.DB) {
    sqlStatement := `DELETE FROM users WHERE id = $1;`
    
    var id int

    fmt.Println("Enter the user id which want to delete : ")
    fmt.Scan(&id)
    _, err := db.Exec(sqlStatement, id)
    if err != nil {
      panic(err)
    }
}