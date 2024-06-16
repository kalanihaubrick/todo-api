package services

import (
	"errors"

	"github.com/kalanihaubrick/todo-api/internal/models"
	"github.com/kalanihaubrick/todo-api/internal/repositories"
)

type TodoService interface {
	GetAllTodos() ([]models.Todo, error)
	GetTodoById(id uint) (*models.Todo, error)
	CreateTodo(todo *models.Todo) error
	UpdateTodo(todo *models.Todo) error
	DeleteTodo(id uint) error
}

type todoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoService{repo}
}

func (s *todoService) GetAllTodos() ([]models.Todo, error) {
	return s.repo.GetAll()
}

func (s *todoService) GetTodoById(id uint) (*models.Todo, error) {
	return s.repo.GetById(id)
}

func (s *todoService) CreateTodo(todo *models.Todo) error {

	if err := s.repo.FindDescription(todo.Task); err == nil {
		return errors.New("a todo with the same description already exists")
	}

	return s.repo.Create(todo)
}

func (s *todoService) UpdateTodo(todo *models.Todo) error {
	_, err := s.GetTodoById(todo.ID)
	if err != nil {
		return err
	}

	if err := s.repo.FindDescription(todo.Task); err == nil {
		return errors.New("a todo with the same description already exists")
	}
	return s.repo.Update(todo)
}

func (s *todoService) DeleteTodo(id uint) error {
	return s.repo.Delete(id)
}
