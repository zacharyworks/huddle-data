package sql

import (
	// SQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/zacharyworks/huddle-shared/data"
	"github.com/zacharyworks/huddle-shared/db"
)

// SelectSessionByID selects a session by its id
func SelectSessionByID(id string) (*types.Session, error) {
	var session types.Session

	rows, err := db.DbCon.Query("SELECT * FROM session WHERE sessionID = ?", id)
	if err != nil {
		return &session, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&session.SessionID,
			&session.State,
			&session.UserFK,
			&session.CreatedDateTime)
		if err != nil {
			return &session, err
		}
	}
	return &session, err
}

// SelectSessionByState selects a session by a state
func SelectSessionByState(state string) (*types.Session, error) {
	var session types.Session

	rows, err := db.DbCon.Query("SELECT * FROM session WHERE state = ?", state)
	if err != nil {
		return &session, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&session.SessionID,
			&session.State,
			&session.UserFK,
			&session.CreatedDateTime)
		if err != nil {
			return &session, err
		}
	}
	return &session, err
}

// InsertSession inserts a session of ID and State
func InsertSession(session types.Session) error {
	stmt, err := db.DbCon.Prepare("INSERT session SET sessionID = ?, state = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.SessionID, session.State)

	if err != nil {
		return err
	}

	return nil
}

// UpdateSession updates a session with new data
func UpdateSession(session types.Session) error {
	_, err := db.DbCon.Query(`
		UPDATE session SET
		sessionID = ?,
		state = ?,
		userFK = ?
		WHERE sessionID = ?`,
		session.SessionID,
		session.State,
		session.UserFK,
		session.SessionID)

	return err
}

// DeleteSession deletes a session by ID
func DeleteSession(session types.Session) error {
	_, err := db.DbCon.Query("DELETE FROM session WHERE sessionID = ?", session.SessionID)
	if err != nil {
		return err
	}
	return nil
}
