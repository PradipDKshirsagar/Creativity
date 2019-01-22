package main 

import (
	"fmt"
	"Creativity/db"
	//"Creativity/user"
	"Creativity/interest"
)

func main() {
    db := db.DBConnetion()
    defer db.Close()
    fmt.Printf("%T",db)

   	//user.Create(db)
    //user.Read(db)
    //user.Update(db)
    //user.Delete(db)

    interest.Create(db)
    interest.Read(db)
    interest.Update(db)
    interest.Delete(db)
}	