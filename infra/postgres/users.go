package postgres

import (
	"database/sql"
	"gqlgen-postgres-demo/usecase"
)

type UserGetter struct {
	db *sql.DB
}

func NewUserGetter(db *sql.DB) *UserGetter {
	return &UserGetter{db}
}

func (u *UserGetter) GetUsers() ([]usecase.UserColumns, error) {
	var us []usecase.UserColumns
	rows, err := u.db.Query("SELECT * FROM users u")
	if err != nil {
		return us, err
	}
	for rows.Next() {
		var u usecase.UserColumns
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return us, err
		}
		us = append(us, u)
	}
	return us, nil
}
