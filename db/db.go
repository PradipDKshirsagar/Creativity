package db

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq" 
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pradip"
	dbname   = "UserInterestApp"
)

var Db *sql.DB

func DBConnetion() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    fmt.Println(db)
    if err != nil {
       fmt.Println("Unsuccessful to connect")
       panic(err)
    }
    err = db.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println("Successfully connected!")
    Db = db
}