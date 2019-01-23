package user

import (
	"creativity/db"
	"database/sql"
	"errors"
	"fmt"
)

func createService(data User) (error){
	db := db.Db 
    sqlStatement := `INSERT INTO users (age, first_name, last_name) VALUES ($1, $2, $3) RETURNING id`
    id := 0
    err := db.QueryRow(sqlStatement, data.Age, data.FirstName, data.LastName).Scan(&id)
    if err != nil {
        return err
    }
    return nil
}

func readService(id string) (User,error){
	db := db.Db

    sqlStatement := `SELECT * FROM USERS WHERE ID = $1;`
    var ur User
    row := db.QueryRow(sqlStatement, id)

    err := row.Scan(&ur.ID, &ur.Age, &ur.FirstName,&ur.LastName)
    switch err {
    case sql.ErrNoRows:
        return User{}, errors.New("user id not exist")
    case nil:
        return ur,nil
    default:
        return User{}, err
    }
}

func updateService(id string,data User) (error){
    db := db.Db

    sqlStatement := `UPDATE users SET age = $2, first_name = $3, last_name = $4 WHERE id = $1;`
    res, err := db.Exec(sqlStatement, id, data.Age, data.FirstName, data.LastName)
    if err != nil {
        return errors.New("Some query fault")
    }
    count, err := res.RowsAffected()
    if err != nil {
        return errors.New("Some query fault")
    }
    if count == 0 {
        return errors.New("Invalid user Id... ")
    }
    fmt.Println(count)
    return nil
}

func deleteService(id string) error {
	db := db.Db

	sqlStatement := `DELETE FROM users_interests WHERE user_id = $1;`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

    sqlStatement = `DELETE FROM users WHERE id = $1;`

    res, err := db.Exec(sqlStatement, id)
    if err != nil {
    	return err
    }
    cnt, err := res.RowsAffected()
    if err != nil {
    	return errors.New("Internal Errors")
    }
    if cnt == 0 {
    	return errors.New("Invalid user id")
    }
    return nil
}

func readAllService()([]User, error){
	db := db.Db 
	var users []User
    sqlStatement := `SELECT * FROM USERS;`
    rows, err := db.Query(sqlStatement)
    if err != nil {
        return users,errors.New("Err to retrive users")
    }

    defer rows.Close()
    var ur User
    for rows.Next() {
        err = rows.Scan(&ur.ID, &ur.Age, &ur.FirstName,&ur.LastName)
        if err != nil {
            return users,errors.New("Err to retrive users")
        }
      //  fmt.Println(ur)
        users = append(users,ur)
    }
    fmt.Println(users)
    return users,nil
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
