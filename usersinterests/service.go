package usersinterests

import (
	"creativity/db"
	"creativity/interest"
	"creativity/user"
	"errors"
	"fmt"
)

func checkService(id string) ([]interest.Interest, error) {
	db := db.Db
	sqlStatement := `SELECT * FROM interests where Id in (select interest_id from users_interests where user_id = $1);`

	var usIn []interest.Interest
	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		return []interest.Interest{}, errors.New("Internal Error ...")
		// fmt.Fprint(w,"No interests present")
	}
	defer rows.Close()
	var in interest.Interest
	for rows.Next() {
		err = rows.Scan(&in.ID, &in.Name)
		if err != nil {
			return []interest.Interest{}, errors.New("Internal Error...")
		}
		fmt.Println(in)
		usIn = append(usIn, in)
	}
	fmt.Println(usIn)
	return usIn, nil
}

func addService(data usersInterests) error {
	db := db.Db

	if !user.Check(data.UsID) {
		return errors.New("Invalid user id")
	}
	if !interest.Check(data.InID) {
		return errors.New("Invalid interest id")
	}

	sqlStatement := `INSERT INTO users_interests (user_id, interest_id) VALUES ($1, $2) RETURNING id`

	id := 0
	err := db.QueryRow(sqlStatement, data.UsID, data.InID).Scan(&id)
	if err != nil {
		panic(err)
	}
	return nil
}
