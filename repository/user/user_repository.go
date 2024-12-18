package repository

import (
	"database/sql"
	"echo-golang/model"
	model_request "echo-golang/model/request"
	model_response "echo-golang/model/response"
)

type IUserRepository interface {
	GetUsers() ([]model_response.User, model.Metadata, error)
	GetUser(id int) (model_response.User, error)
	LoginUser(login model_request.Login) (model_response.User, error)
	RegisterUser(register model_request.Register) (sql.Result, error)
}
