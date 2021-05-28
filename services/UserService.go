package services

import (
	"RestApi/db"
	"RestApi/models"
	"database/sql"
	"fmt"
	"log"
)


func InsertUser(user models.User) int64{
 connection:= db.InitDB() // create the postgres db connection
	defer connection.Close() // close connection
	sqlStatement := `INSERT into users (name,location, age) VALUES ($1, $2, $3) `
	var status int64  = 1
	_, err := connection.Exec(sqlStatement, user.Name, user.Location, user.Age)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
		return 0
	}
	return status
}

func GetUser(id int64) (models.User, error){
	connection:= db.InitDB() // create the postgres db connection
	defer connection.Close() // close connection
	var user models.User
	sqlStatement := `SELECT * FROM users WHERE id=$1`
	row := connection.QueryRow(sqlStatement, id)
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return user, err
}

// get one user from the DB by its userid
func GetAllUsers() ([]models.User, error) {
	connection:= db.InitDB() // create the postgres db connection
	defer connection.Close() // close connection
	var users []models.User
	// create the select sql query
	sqlStatement := `SELECT * FROM users`
	// execute the sql statement
	rows, err := connection.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	// close the statement
	defer rows.Close()
	// iterate over the rows
	for rows.Next() {
		var user models.User
		// unmarshal the row object to user
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		// append the user in the users slice
		users = append(users, user)
	}
	// return empty user on error
	return users, err
}

// update user in the DB
func UpdateUser(id int64, user models.User) int64 {
	connection:= db.InitDB() // create the postgres db connection
	defer connection.Close() // close connection
	// create the update sql query
	sqlStatement := `UPDATE users SET name=$2, location=$3, age=$4 WHERE id=$1`
	// execute the sql statement
	res, err := connection.Exec(sqlStatement, id, user.Name, user.Location, user.Age)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}

// delete user in the DB
func DeleteUser(id int64) int64 {
	connection:= db.InitDB() // create the postgres db connection
	defer connection.Close() // close connection
	// create the delete sql query
	sqlStatement := `DELETE FROM users WHERE id=$1`
	// execute the sql statement
	res, err := connection.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}