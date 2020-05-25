package postgres

import "database/sql"

type UserColumns struct {
	ID   int
	Name string
}

func QueryUsers(db *sql.DB) ([]UserColumns, error) {
	var us []UserColumns
	rows, err := db.Query("SELECT * FROM users u")
	if err != nil {
		return us, err
	}
	for rows.Next() {
		var u UserColumns
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return us, err
		}
		us = append(us, u)
	}
	return us, nil
}
