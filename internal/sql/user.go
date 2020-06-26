package sql

import (
	// SQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/zacbriggssagecom/huddle/server/sharedinternal/data"
	"github.com/zacbriggssagecom/huddle/server/sharedinternal/db"
	"log"
)

// SelectUser selects a user by their authID
func SelectUser(oauthID string) (*types.User, error) {
	var user types.User
	rows, err := db.DbCon.Query("SELECT * FROM user WHERE oauthID = ?", oauthID)
	if err != nil {
		return &user, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.UserID, &user.OauthID, &user.Email, &user.Picture)
		if err != nil {
			return &user, err
		}
	}
	return &user, err
}

// SelectUserByID selects a user by their sql ID
func SelectUserByID(id int) (*types.User, error) {
	var user types.User
	rows, err := db.DbCon.Query("SELECT * FROM user WHERE userID = ?", id)
	if err != nil {
		return &user, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.UserID, &user.OauthID, &user.Email, &user.Picture)
		if err != nil {
			return &user, err
		}
	}
	return &user, err
}

//InsertUser creates a new user
func InsertUser(user types.User) (int, error) {
	stmt, err := db.DbCon.Prepare("INSERT user SET oauthID = ?, email = ?, picture = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.OauthID, user.Email, user.Picture)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}
