package sql

import (
	// SQL driver
	"../shared"
	_ "github.com/go-sql-driver/mysql"
	types "github.com/zacharyworks/huddle-shared/data"
	"github.com/zacharyworks/huddle-shared/db"
)

func InsertBoardJoinCode(boardJoinCode types.BoardJoinCode) *shared.AppError {
	if _, err := db.DbCon.Query(`INSERT boardJoinCode
			SET boardFK = ?,
			    code = ?`,
			    boardJoinCode.BoardFK,
			    boardJoinCode.Code); err != nil {
		return shared.ErrorInsertingRecord(err)
	}
	return nil
}

func JoinBoard(boardJoin types.BoardJoin) (types.Board, *shared.AppError) {
	var board types.Board
	err := db.DbCon.QueryRow(`
		SELECT board.boardID, 
		       board.Name, 
		       board.boardType  
		FROM board 
    		INNER JOIN boardJoinCode ON board.boardID = boardJoinCode.boardFK 
			WHERE boardJoinCode.code = ?`, boardJoin.Code).Scan(
			&board.BoardID,
			&board.Name,
			&board.BoardType)
	if err != nil {
		return board, shared.ErrorRetrievingRecord(err)
	}

	e := InsertBoardMember(types.BoardMember{
			BoardFK: board.BoardID,
			UserFK: boardJoin.UserFK})
	if e != nil {
		return board, e
	}

	return board, nil
}