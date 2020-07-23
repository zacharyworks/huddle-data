package sql

import (
	// SQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/zacharyworks/huddle-shared/data"
	"github.com/zacharyworks/huddle-shared/db"
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
		err := rows.Scan(&user.OauthID, &user.Email, &user.Picture, &user.Name, &user.GivenName, &user.FamilyName)
		if err != nil {
			return &user, err
		}
	}
	return &user, err
}


//InsertUser creates a new user
func InsertUser(user types.User) (*types.User, error) {
	stmt, err := db.DbCon.Prepare(`INSERT user SET oauthID = ?, 
                email = ?, 
                picture = ?, 
                name = ?, 
                givenName = ?, 
                familyName = ?`)
	if err != nil {
		return &types.User{}, err
	}
	defer stmt.Close()

	// Create user
	_, err = stmt.Exec(user.OauthID, user.Email, user.Picture, user.Name, user.GivenName, user.FamilyName)

	// Create the users personal board
	userBoard, err := InsertBoard(types.Board{Name: user.Name, BoardType: 0})
	if err != nil {
		return &types.User{}, err
	}
	err = InsertBoardMember(types.BoardMember{UserFK: user.OauthID, BoardFK: userBoard.BoardID})
	if err != nil {
		return &types.User{}, err
	}

	return SelectUser(user.OauthID)
}
