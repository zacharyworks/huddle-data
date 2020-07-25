package sql

import (
	"../shared"
	"net/http"

	// SQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/zacharyworks/huddle-shared/data"
	"github.com/zacharyworks/huddle-shared/db"
)

// SelectUser selects a user by their authID
func SelectUser(oauthID string) (*types.User, *shared.AppError) {
	var user types.User
	if err := db.DbCon.QueryRow("SELECT * FROM user WHERE oauthID = ?", oauthID).Scan(
		&user.OauthID,
		&user.Email,
		&user.Picture,
		&user.Name,
		&user.GivenName,
		&user.FamilyName); err != nil {
		return &user, &shared.AppError{err, "User not found", http.StatusNotFound}
	}

	return &user, nil
}


//InsertUser creates a new user
func InsertUser(user types.User) (*types.User, *shared.AppError) {
	stmt, err := db.DbCon.Prepare(`INSERT user SET oauthID = ?, 
                email = ?, 
                picture = ?, 
                name = ?, 
                givenName = ?, 
                familyName = ?`)
	defer stmt.Close()
	if err != nil {
		return nil, shared.ErrorRetrievingRecord(err)
	}


	// Create user
	if _, err = stmt.Exec(
		user.OauthID,
		user.Email,
		user.Picture,
		user.Name,
		user.GivenName,
		user.FamilyName); err != nil {
		return nil, shared.ErrorInsertingRecord(err)
	}

	// Create the users personal board
	userBoard, e := InsertBoard(types.Board{Name: user.Name, BoardType: 0})
	if e != nil {
		return nil, e
	}

	if e := InsertBoardMember(types.BoardMember{UserFK: user.OauthID, BoardFK: userBoard.BoardID}); e != nil {
		return nil, e
	}

	return SelectUser(user.OauthID)
}
