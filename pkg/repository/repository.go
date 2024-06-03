package repository

import (
	"Resrik"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user Resrik.User) (int, error)
	GetUser(username, password string) (Resrik.User, error)
}

type TodoList interface {
	Create(userId int, list Resrik.TodoList) (int, error)
	GetAll(userId int) ([]Resrik.TodoList, error)
	GetById(userId int, listId int) (Resrik.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input Resrik.UpdateInput) error
}

type TodoItem interface {
	Create(listId int, item Resrik.TodoItem) (int, error)
	GetAll(userId, listId int) ([]Resrik.TodoItem, error)
	GetByItemId(userId, itemId int) ([]Resrik.TodoItem, error)
	Delete(userId, itemId int) error
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
