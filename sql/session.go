package sql

import (
	"github.com/zacharyworks/huddle-data/shared"
	"net/http"

	// SQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/zacharyworks/huddle-shared/data"
	"github.com/zacharyworks/huddle-shared/db"
)

// SelectSessionByID selects a session by its id
func SelectSessionByID(id string) (*types.Session, *shared.AppError) {
	var session types.Session

	if err := db.DbCon.QueryRow("SELECT * FROM session WHERE sessionID = ?", id).Scan(
		&session.SessionID,
		&session.State,
		&session.UserFK,
		&session.CreatedDateTime); err != nil {
		return &types.Session{}, &shared.AppError{
			err,
			"Session not found",
			http.StatusNotFound}
	}

	return &session, nil
}

// SelectSessionByState selects a session by a state
func SelectSessionByState(state string) (*types.Session, *shared.AppError) {
	var session types.Session

	if err := db.DbCon.QueryRow("SELECT * FROM session WHERE state = ?", state).Scan(
		&session.SessionID,
		&session.State,
		&session.UserFK,
		&session.CreatedDateTime); err != nil {
		return nil, shared.ErrorRetrievingRecord(err)
	}

	return &session, nil
}

// InsertSession inserts a session of ID and State
func InsertSession(session types.Session) *shared.AppError {
	stmt, err := db.DbCon.Prepare("INSERT session SET sessionID = ?, state = ?")
	if err != nil {
		return shared.ErrorInsertingRecord(err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(session.SessionID, session.State); err != nil {
		return shared.ErrorInsertingRecord(err)
	}

	return nil
}

// UpdateSession updates a session with new data
func UpdateSession(session types.Session) *shared.AppError {
	if _, err := db.DbCon.Query(`
		UPDATE session SET
		sessionID = ?,
		state = ?,
		userFK = ?
		WHERE sessionID = ?`,
		session.SessionID,
		session.State,
		session.UserFK,
		session.SessionID); err != nil {
		return shared.ErrorUpdatingRecord(err)
	}

	return nil
}

// DeleteSession deletes a session by ID
func DeleteSession(session types.Session) *shared.AppError {
	_, err := db.DbCon.Query("DELETE FROM session WHERE sessionID = ?", session.SessionID)
	if err != nil {
		return shared.ErrorDeletingRecord(err)
	}
	return nil
}
