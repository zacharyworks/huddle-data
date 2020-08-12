package sql

import (
	"github.com/zacharyworks/huddle-data/shared"
	// SQL driver
	_ "github.com/go-sql-driver/mysql"
	types "github.com/zacharyworks/huddle-shared/data"
	"github.com/zacharyworks/huddle-shared/db"
)

// SelectBoardByID selects a board by its id
func SelectBoardMemberByMemberID(id string) (board *types.Board, e *shared.AppError) {
	board = &types.Board{}
	if err := db.DbCon.QueryRow("SELECT * FROM boardMember WHERE userFK = ?", id).Scan(
		&board.BoardID,
		&board.Name); err != nil {
		return board, shared.ErrorRetrievingRecord(err)
	}

	return board, nil
}

// InsertBoardMember inserts a new member for a board
func InsertBoardMember(boardMember types.BoardMember) *shared.AppError {
	if _, err := db.DbCon.Exec(`INSERT boardMember 
			SET boardFK = ?, 
			    userFK = ?`,
			boardMember.BoardFK,
			boardMember.UserFK); err != nil {
		return shared.ErrorInsertingRecord(err)
	}
	return nil
}

// UpdateBoardMember updates a Board Member entry with new data
func UpdateBoardMember(boardMember types.BoardMember) *shared.AppError {
	if _, err := db.DbCon.Exec(`
		UPDATE boardMember SET
		boardFK = ?,
		userFK = ?
		WHERE boardFK = ? AND userFK = ?`,
		boardMember.BoardFK,
		boardMember.UserFK,
		boardMember.BoardFK,
		boardMember.UserFK); err != nil {
		return shared.ErrorUpdatingRecord(err)
	}
	return nil
}

// DeleteBoardMember deletes a board member given IDs
func DeleteBoardMember(boardMember types.BoardMember) *shared.AppError {
	if _, err := db.DbCon.Exec(`
		DELETE FROM boardMember 
		WHERE boardFK = ? 
		AND userFK = ?`,
		boardMember.BoardFK,
		boardMember.UserFK); err != nil {
		return shared.ErrorDeletingRecord(err)
	}
	return nil
}
