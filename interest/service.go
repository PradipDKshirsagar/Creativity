package interest

import (
	"github.com/PradipDKshirsagar/Creativity/db"
	"database/sql"
	"errors"
	"fmt"
)

func createService(data Interest) error {
	db := db.Db
	sqlStatement := `INSERT INTO interests (name) VALUES ($1) RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, data.Name).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func readService(id string) (Interest, error) {
	db := db.Db

	sqlStatement := `SELECT * FROM interests WHERE ID = $1`
	var in Interest
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&in.ID, &in.Name)
	switch err {
	case sql.ErrNoRows:
		return Interest{}, errors.New("Interest id not exist")
	case nil:
		return in, nil
	default:
		return Interest{}, err
	}
}

func updateService(id string, data Interest) error {
	db := db.Db

	sqlStatement := `UPDATE interests SET name = $2 WHERE id = $1;`
	res, err := db.Exec(sqlStatement, id, data.Name)
	if err != nil {
		return errors.New("Some query fault")
	}
	count, err := res.RowsAffected()
	if err != nil {
		return errors.New("Some query fault")
	}
	if count == 0 {
		return errors.New("Invalid Interest Id... ")
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

	sqlStatement = `DELETE FROM interests WHERE id = $1;`

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	cnt, err := res.RowsAffected()
	if err != nil {
		return errors.New("Internal Errors")
	}
	if cnt == 0 {
		return errors.New("Invalid Interest id")
	}
	return nil
}

func readAllService() ([]Interest, error) {
	db := db.Db
	var interests []Interest
	sqlStatement := `SELECT * FROM interests;`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return interests, errors.New("Err to retrive interests")
	}

	defer rows.Close()
	var in Interest
	for rows.Next() {
		err = rows.Scan(&in.ID, &in.Name)
		if err != nil {
			return interests, errors.New("Err to retrive interests")
		}
		//  fmt.Println(ur)
		interests = append(interests, in)
	}
	fmt.Println(interests)
	return interests, nil
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
