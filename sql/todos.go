package sql

import
(
	"github.com/zacharyworks/huddle-data/shared"
	// SQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/zacharyworks/huddle-shared/data"
	"github.com/zacharyworks/huddle-shared/db"
)

// GetTodosForBoard gets to-dos for a specific board
func GetTodosForBoard(id int) (todos []types.Todo, e *shared.AppError) {
	rows, err := db.DbCon.Query(`SELECT * FROM todo WHERE boardFK = ?`, id)
	if err != nil {
		return nil, shared.ErrorRetrievingRecord(err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo types.Todo
		if err = rows.Scan(&todo.TodoID, &todo.Status, &todo.Value, &todo.ParentFK, &todo.BoardFK); err != nil {
			return nil, shared.ErrorParsingRecord(err)
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

// SelectTodo selects a to-do based off its ID
func SelectTodo(id int) (*types.Todo, *shared.AppError) {
	var todo types.Todo

	if err := db.DbCon.QueryRow("SELECT * FROM todo WHERE todoID = ?", id).Scan(
		&todo.TodoID,
		&todo.Status,
		&todo.Value,
		&todo.ParentFK,
		&todo.BoardFK); err != nil {
		return &todo, shared.ErrorRetrievingRecord(err)
	}

	return &todo, nil
}

// InsertTodo inserts a new to-do
func InsertTodo(todo types.Todo) (int, *shared.AppError) {
	stmt, err := db.DbCon.Prepare("INSERT todo SET status = ?, value = ?, parentFK = ?, boardFK = ?")
	defer stmt.Close()
	if err != nil {
		return 0, shared.ErrorInsertingRecord(err)
	}

	result, err := stmt.Exec(todo.Status, todo.Value, todo.ParentFK, todo.BoardFK)
	if err != nil {
		return 0, shared.ErrorParsingRecord(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, shared.ErrorRetrievingRecord(err)
	}
	return int(id), nil
}

// UpdateTodo updates a to-do
func UpdateTodo(todo types.Todo) *shared.AppError {
	if _, err := db.DbCon.Exec("UPDATE todo SET status = ?, value = ?, parentFK = ? WHERE todoID = ?",
		todo.Status,
		todo.Value,
		todo.ParentFK,
		todo.TodoID); err != nil {
		return shared.ErrorUpdatingRecord(err)
	}

	return nil
}

// RemoveTodo removes a to-do from the database
func RemoveTodo(todo types.Todo) *shared.AppError {
	_, err := db.DbCon.Exec("DELETE FROM todo WHERE todoID = ?", todo.TodoID)
	if err != nil {
		return shared.ErrorProcessingParameter(err)
	}

	return nil
}
