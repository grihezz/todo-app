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
}

type TodoItem interface {
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
	}
}
