package sql

import (
	// SQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/zacharyworks/huddle-shared/data"
	"github.com/zacharyworks/huddle-shared/db"
)

func SelectUsersBoards(id string) (boards []types.Board, err error) {
	var (
		boardFK int
		userFK string
		)
	rows, err := db.DbCon.Query(
	`SELECT * FROM board 
			INNER JOIN boardMember ON board.boardID = boardMember.boardFK
			WHERE boardMember.userFK = ?`, id)

	defer rows.Close()
	for rows.Next() {
		var newBoard types.Board
		err := rows.Scan(&newBoard.BoardID, &newBoard.Name, &newBoard.BoardType, &boardFK, &userFK)
		if err != nil {
			return boards, err
		}
		boards = append(boards, newBoard)
	}

	return boards, err
}

// SelectBoardByID selects a board by its id
func SelectBoardByID(id int) (*types.Board, error) {
	rows, err := db.DbCon.Query("SELECT * FROM board WHERE boardID = ?", id)
	var board types.Board
	if err != nil {
		return &board, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&board.BoardID,
			&board.Name,
			&board.BoardType)
		if err != nil {
			return &board, err
		}
	}
	return &board, err
}

// InsertBoard inserts a board of ID and Name
func InsertBoard(board types.Board) (*types.Board, error) {
	stmt, err := db.DbCon.Prepare("INSERT board SET name = ?, boardType = ?")
	if err != nil {
		return &types.Board{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(board.Name, board.BoardType)

	if err != nil {
		return &types.Board{}, err
	}

	id, err := result.LastInsertId()
	return SelectBoardByID(int(id))
}

// UpdateBoard updates a session with new data
func UpdateBoard(board types.Board) error {
	_, err := db.DbCon.Query(`
		UPDATE board SET
		boardID = ?,
		name = ?,
	 	boardType = ?
		WHERE boardID = ?`,
		board.BoardID,
		board.Name,
		board.BoardID,
		board.BoardType)

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
