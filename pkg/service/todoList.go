package service

import (
	"Resrik"
	"Resrik/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (s *TodoListService) Create(userId int, list Resrik.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]Resrik.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId int, listId int) (Resrik.TodoList, error) {
	return s.repo.GetById(userId, listId)
}
