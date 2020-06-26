package sql

import (
	// SQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/zacbriggssagecom/huddle/server/sharedinternal/data"
	"github.com/zacbriggssagecom/huddle/server/sharedinternal/db"
)

// SelectBoardByID selects a board by its id
func SelectBoardMemberByMemberID(id string) (board *types.Board, err error) {
	rows, err := db.DbCon.Query("SELECT * FROM boardMember WHERE userFK = ?", id)
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

// InsertBoardMember inserts a new member for a board
func InsertBoardMember(boardMember types.BoardMember) error {
	stmt, err := db.DbCon.Prepare("INSERT boardMember SET boardFK = ?, userFK = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(boardMember.BoardFK, boardMember.UserFK)

	if err != nil {
		return err
	}

	return nil
}

// UpdateBoardMember updates a Board Member entry with new data
func UpdateBoardMember(boardMember types.BoardMember) error {
	_, err := db.DbCon.Query(`
		UPDATE board SET
		boardFK = ?,
		userFK = ?
		WHERE boardID = ? AND userFK = ?`,
		boardMember.BoardFK,
		boardMember.UserFK,
		boardMember.BoardFK,
		boardMember.UserFK)

	return err
}

// DeleteBoardMember deletes a boardmember given IDs
func DeleteBoardMember(boardMember types.BoardMember) error {
	_, err := db.DbCon.Query("DELETE FROM boardMember WHERE boardFK = ? AND userFK = ?", boardMember.BoardFK, boardMember.UserFK)
	if err != nil {
		return err
	}
	return nil
}
