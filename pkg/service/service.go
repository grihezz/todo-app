package service

import (
	"Resrik"
	"Resrik/pkg/repository"
)

type Authorization interface {
	CreateUser(user Resrik.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list Resrik.TodoList) (int, error)
	GetAll(userId int) ([]Resrik.TodoList, error)
	GetById(userId int, listId int) (Resrik.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId, listId int, input Resrik.UpdateInput) error
}

type TodoItem interface {
	Create(userId int, listId int, item Resrik.TodoItem) (int, error)
	GetAll(userId, listId int) ([]Resrik.TodoItem, error)
	GetByItemId(userId int, itemId int) ([]Resrik.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, id int, input Resrik.UpdateInputItem) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
		TodoItem:      NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}
