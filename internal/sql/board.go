package sql

import (
	// SQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/zacbriggssagecom/huddle/server/sharedinternal/data"
	"github.com/zacbriggssagecom/huddle/server/sharedinternal/db"
)

// SelectBoardByID selects a board by its id
func SelectBoardByID(id string) (board *types.Board, err error) {
	rows, err := db.DbCon.Query("SELECT * FROM session WHERE sessionID = ?", id)
	if err != nil {
		return board, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&board.BoardID,
			&board.Name)
		if err != nil {
			return board, err
		}
	}
	return board, err
}

// InsertBoard inserts a board of ID and Name
func InsertBoard(board types.Board) error {
	stmt, err := db.DbCon.Prepare("INSERT board SET boardID = ?, name = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(board.BoardID, board.Name)

	if err != nil {
		return err
	}

	return nil
}

// UpdateBoard updates a session with new data
func UpdateBoard(board types.Board) error {
	_, err := db.DbCon.Query(`
		UPDATE board SET
		boardID = ?,
		name = ?
		WHERE boardID = ?`,
		board.BoardID,
		board.Name,
		board.BoardID)

	return err
}

// DeleteBoard deletes a session by ID
func DeleteBoard(board types.Board) error {
	_, err := db.DbCon.Query("DELETE FROM board WHERE boardID = ?", board.BoardID)
	if err != nil {
		return err
	}
	return nil
}
