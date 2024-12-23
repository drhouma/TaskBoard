package postgres

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
)

const (
	tasksTable = "tasks"
)

type Task struct {
	Id          int32   `json:"id,omitempty" db:"id" goqu:"skipinsert"`
	User        *string `json:"user" db:"user"`
	Description string  `json:"description" db:"description"`
	Category    *string `json:"category" db:"category"`
}

func (db *DataBase) AddTask(item Task) error {
	insertQuery, _, err := goqu.Insert(tasksTable).Rows(item).ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}
	if _, err = db.DB.Exec(insertQuery); err != nil {
		return fmt.Errorf("insert data: %w", err)
	}
	return nil
}

func (db *DataBase) GetTasks() ([]Task, error) {
	selectQuery, _, err := goqu.From(tasksTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var tasks []Task
	if err = db.DB.Select(&tasks, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}
	return tasks, nil
}

func (db *DataBase) DeleteTask(user string, description string) error {
	deleteQuery, _, err := goqu.Delete(tasksTable).Where(goqu.And(
		goqu.C("user").Eq(user),
		goqu.C("description").Eq(description))).
		Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}

	var id string
	row := db.DB.QueryRowx(deleteQuery)

	if err = row.Scan(&id); err != nil {
		return fmt.Errorf("delete data: %w", err)
	}
	return nil
}

func (db *DataBase) UpdateTask(user string, description string, category string) error {
	updateQuery, _, err := goqu.Update(tasksTable).Set(goqu.Record{"category": category}).Where(goqu.And(
		goqu.C("user").Eq(user),
		goqu.C("description").Eq(description))).
		Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}

	var id string
	row := db.DB.QueryRowx(updateQuery)

	if err = row.Scan(&id); err != nil {
		return fmt.Errorf("delete data: %w", err)
	}
	return nil
}
