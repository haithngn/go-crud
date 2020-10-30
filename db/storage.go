package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/haithngn/go-crud/model"
)

type Storage struct {
	Database *sql.DB
}

func (storage *Storage) Create(params struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}) (*model.Question, error) {
	state, err := storage.Database.Prepare("insert into questions (title, content) values (?, ?)")
	if err != nil {
		return nil, err
	}
	rows, err := state.Exec(params.Title, params.Content)
	if err != nil {
		return nil, err
	}
	id, err := rows.LastInsertId()

	question := model.Question{
		ID:      int(id),
		Title:   params.Title,
		Content: params.Content,
	}
	return &question, nil
}

func (storage *Storage) Find(id int) (*model.Question, error) {
	var question model.Question
	err := storage.Database.QueryRow("select id, title, content from questions where id = ?", id).Scan(&question.ID, &question.Title, &question.Content)
	if err != nil {
		return nil, err
	}

	return &question, nil
}
