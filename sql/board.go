package sql

import (
	"github.com/zacharyworks/huddle-data/shared"
	// SQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/zacharyworks/huddle-shared/data"
	"github.com/zacharyworks/huddle-shared/db"
)

func SelectUsersBoards(id string) (boards []types.Board, e *shared.AppError) {
	rows, err := db.DbCon.Query(
	`SELECT board.BoardID, board.Name, board.BoardType FROM board 
			INNER JOIN boardMember ON board.boardID = boardMember.boardFK
			WHERE boardMember.userFK = ?`, id)
	if err != nil {
		return boards, shared.ErrorRetrievingRecord(err)
	}

	defer rows.Close()

	for rows.Next() {
		var newBoard types.Board
		err := rows.Scan(&newBoard.BoardID, &newBoard.Name, &newBoard.BoardType)
		if err != nil {
			return boards, shared.ErrorParsingRecord(err)
		}
		boards = append(boards, newBoard)
	}

	return boards, nil
}

// SelectBoardByID selects a board by its id
func SelectBoardByID(id int) (board *types.Board, e *shared.AppError) {
	board = &types.Board{}
	err := db.DbCon.QueryRow("SELECT * FROM board WHERE boardID = ?", id).Scan(
		&board.BoardID,
		&board.Name,
		&board.BoardType)

	if err != nil {
		return nil, shared.ErrorRetrievingRecord(err)
	}
	return
}

// InsertBoard inserts a board of ID and Name
func InsertBoard(board types.Board) (*types.Board, *shared.AppError) {
	stmt, err := db.DbCon.Prepare("INSERT board SET name = ?, boardType = ?")
	if err != nil {
		return &types.Board{}, shared.ErrorInsertingRecord(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(board.Name, board.BoardType)

	if err != nil {
		return nil, shared.ErrorInsertingRecord(err)
	}

	if id, err := result.LastInsertId(); err != nil {
		return nil, shared.ErrorRetrievingRecord(err)
	} else {
		return SelectBoardByID(int(id))
	}
}

// UpdateBoard updates a session with new data
func UpdateBoard(board types.Board) *shared.AppError {
	_, err := db.DbCon.Exec(`
		UPDATE board SET
		boardID = ?,
		name = ?,
	 	boardType = ?
		WHERE boardID = ?`,
		board.BoardID,
		board.Name,
		board.BoardID,
		board.BoardType)
	if err != nil {
		return shared.ErrorUpdatingRecord(err)
	}
	return nil
}

// DeleteBoard deletes a session by ID
func DeleteBoard(board types.Board) *shared.AppError {
	_, err := db.DbCon.Exec("DELETE FROM board WHERE boardID = ?", board.BoardID)
	if err != nil {
		return shared.ErrorDeletingRecord(err)
	}
	return nil
}
