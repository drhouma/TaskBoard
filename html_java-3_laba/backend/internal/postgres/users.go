package postgres

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
)

const (
	usersTable = "users"
)

type User struct {
	Name string `json:"name" db:"name"`
}

func (db *DataBase) AddUser(item User) error {
	insertQuery, _, err := goqu.Insert(usersTable).Rows(item).ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}
	if _, err = db.DB.Exec(insertQuery); err != nil {
		return fmt.Errorf("insert data: %w", err)
	}
	return nil
}

func (db *DataBase) IsUserExist(item User) (bool, error) {
	selectQuery, _, err := goqu.From(usersTable).Where(goqu.C("name").Eq(item.Name)).ToSQL()
	if err != nil {
		return false, fmt.Errorf("configure query: %w", err)
	}

	var users []User
	if err = db.DB.Select(&users, selectQuery); err != nil {
		return false, fmt.Errorf("select data: %w", err)
	}
	return users != nil, nil
}
