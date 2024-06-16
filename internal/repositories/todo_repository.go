package repositories

import (
	"github.com/kalanihaubrick/todo-api/internal/models"
	"gorm.io/gorm"
)

type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	GetById(id uint) (*models.Todo, error)
	Create(todo *models.Todo) error
	Update(todo *models.Todo) error
	Delete(id uint) error
	FindDescription(description string) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db}
}

func (r *todoRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *todoRepository) GetById(id uint) (*models.Todo, error) {
	var todo models.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *todoRepository) Create(todo *models.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepository) Delete(id uint) error {
	return r.db.Delete(&models.Todo{}, id).Error
}

func (r *todoRepository) Update(todo *models.Todo) error {
	return r.db.Save(todo).Error
}

func (r *todoRepository) FindDescription(description string) error {
	var existingTodo models.Todo
	return r.db.Where("task = ?", description).First(&existingTodo).Error
}
