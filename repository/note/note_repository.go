package repository

import (
	"database/sql"
	"echo-golang/model"
	model_request "echo-golang/model/request"
	model_response "echo-golang/model/response"
)

type INoteRepository interface {
	GetNote(pagination model.Pagination) ([]model_response.Note, error)
	InsertNote(note model_request.Note) (sql.Result, error)
	DeleteNoteById(id int) (sql.Result, error)
	UpdateNoteById(id int, note model_request.Note) (sql.Result, error)
}
