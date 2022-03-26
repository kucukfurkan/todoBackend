package service

import (
	"Desktop/todo-backend/model"
	"Desktop/todo-backend/repository"
)

type Service interface {
	CreateTodo(todo model.SendTodoElements) error
	GetTodoElements() (todos []model.TodoElements, err error)
}

type service struct {
	repo repository.Repository
}

var _ Service = service{}

func NewService(repo repository.Repository) Service {
	return service{repo: repo}
}

func (s service) GetTodoElements() (todos []model.TodoElements, err error) {
	return s.repo.GetTodoElements()
}

func (s service) CreateTodo(todo model.SendTodoElements) error {
	return s.repo.CreateTodo(todo)
}
