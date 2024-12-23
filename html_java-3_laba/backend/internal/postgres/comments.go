package postgres

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
)

const (
	commentsTable = "comments"
)

type Comment struct {
	Id       uint64  `json:"id,omitempty" db:"id" goqu:"skipinsert"`
	UserName *string `json:"user_name" db:"user_name"`
	Message  *string `json:"message" db:"message"`
}

func (db *DataBase) AddComment(item Comment) error {
	insertQuery, _, err := goqu.Insert(commentsTable).Rows(item).ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}
	if _, err = db.DB.Exec(insertQuery); err != nil {
		return fmt.Errorf("insert data: %w", err)
	}
	return nil
}

func (db *DataBase) GetComments() ([]Comment, error) {
	selectQuery, _, err := goqu.From(commentsTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var comments []Comment
	if err = db.DB.Select(&comments, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}
	return comments, nil
}

func (db *DataBase) DeleteComment(item Comment) error {
	deleteQuery, _, err := goqu.Delete(commentsTable).Where(goqu.And(
		goqu.C("user_name").Eq(item.UserName),
		goqu.C("message").Eq(item.Message))).
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

func (db *DataBase) IsCommentExist(item Comment) (bool, error) {
	selectQuery, _, err := goqu.From(commentsTable).Where(goqu.And(
		goqu.C("user_name").Eq(item.UserName),
		goqu.C("message").Eq(item.Message))).ToSQL()
	if err != nil {
		return false, fmt.Errorf("configure query: %w", err)
	}

	var comments []Comment
	if err = db.DB.Select(&comments, selectQuery); err != nil {
		return false, fmt.Errorf("select data: %w", err)
	}
	return comments != nil, nil
}
