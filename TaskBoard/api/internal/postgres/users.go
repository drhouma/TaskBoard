package postgres

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
)

const (
	usersTable = "users"
)

type User struct {
	Name     string `json:"name" db:"name"`
	Password string `json:"password" db:"password"`
}

func (db *DataBase) GetUser(name string) (*User, error) {
	selectQuery, _, err := goqu.From(usersTable).Where(goqu.C("name").Eq(name)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var users []User
	if err = db.DB.Select(&users, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}
	return &users[0], nil
}
