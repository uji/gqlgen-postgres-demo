package postgres

import (
	"database/sql"
	"fmt"
	"gqlgen-postgres-demo/usecase"
)

type TodoGetter struct {
	db *sql.DB
}

func NewTodoGetter(db *sql.DB) *TodoGetter {
	return &TodoGetter{db}
}

func (t *TodoGetter) GetTodosByUserID(userID int) ([]usecase.TodoColumns, error) {
	ts := make([]usecase.TodoColumns, 0, 2)
	sql := fmt.Sprintf("SELECT t.id, t.text, t.done, t.name FROM todos t WHERE user_id = %d", userID)
	rows, err := t.db.Query(sql)
	if err != nil {
		return ts, err
	}
	for rows.Next() {
		var t usecase.TodoColumns
		if err := rows.Scan(&t.ID, &t.Text, &t.Done, &t.Name); err != nil {
			return ts, err
		}
		ts = append(ts, t)
	}
	return ts, nil
}
