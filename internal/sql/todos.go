package sql

import (
	// SQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/zacbriggssagecom/huddle/server/sharedinternal/data"
	"github.com/zacbriggssagecom/huddle/server/sharedinternal/db"
	"log"
)

// GetAllTodo selects ALL todos from the database
func GetAllTodo() (*[]types.Todo, error) {
	var (
		todoID   int
		status   int
		parentFK int
		boardFK  int
		value    string
		allTodo  []types.Todo
	)
	rows, err := db.DbCon.Query("SELECT * FROM todo")
	if err != nil {
		return &allTodo, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&todoID, &status, &value, &parentFK, &boardFK)
		if err != nil {
			return &allTodo, err
		}

		allTodo = append(allTodo, types.Todo{
			TodoID:   todoID,
			Status:   status,
			Value:    value,
			ParentFK: parentFK,
			BoardFK:  boardFK,
		})
	}
	return &allTodo, err
}

// SelectTodo selects a todo based off its ID
func SelectTodo(id int) (*types.Todo, error) {
	var todo types.Todo

	rows, err := db.DbCon.Query("SELECT * FROM todo WHERE todoID = ?", id)
	if err != nil {
		return &todo, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&todo.TodoID, &todo.Status, &todo.Value, &todo.ParentFK, &todo.BoardFK)
		if err != nil {
			return &todo, err
		}

	}
	return &todo, err
}

// InsertTodo inserts a new todo
func InsertTodo(todo types.Todo) (int, error) {
	stmt, err := db.DbCon.Prepare("INSERT todo SET status = ?, value = ?, parentFK = ?, boardFK = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(todo.Status, todo.Value, todo.ParentFK, todo.BoardFK)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

// UpdateTodo updates a todo
func UpdateTodo(todo types.Todo) error {
	_, err := db.DbCon.Query("UPDATE todo SET status = ?, value = ?, parentFK = ? WHERE todoID = ?",
		todo.Status,
		todo.Value,
		todo.ParentFK,
		todo.TodoID)

	return err
}

// RemoveTodo removes a todo from the database
func RemoveTodo(todo types.Todo) error {

	_, err := db.DbCon.Query("DELETE FROM todo WHERE todoID = ?", todo.TodoID)

	return err
}
