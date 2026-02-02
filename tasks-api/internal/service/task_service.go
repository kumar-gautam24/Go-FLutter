package service

import (
	"errors"
	"tasks-api/internal/domain"
	"tasks-api/internal/repository"
)

// TaskService handles business logic (similar to use cases in Flutter)
type TaskService struct {
	repo repository.TaskRepository
}

// NewTaskService creates a new task service
func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

// CreateTask creates a new task
func (s *TaskService) CreateTask(title, description string) (*domain.Task, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	task := domain.NewTask(title, description)
	err := s.repo.Create(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// GetAllTasks retrieves all tasks
func (s *TaskService) GetAllTasks() ([]*domain.Task, error) {
	return s.repo.GetAll()
}

// GetTaskByID retrieves a task by ID
func (s *TaskService) GetTaskByID(id string) (*domain.Task, error) {
	return s.repo.GetByID(id)
}

// ToggleTaskCompletion toggles the completion status of a task
func (s *TaskService) ToggleTaskCompletion(id string) (*domain.Task, error) {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	task.IsCompleted = !task.IsCompleted
	err = s.repo.Update(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// DeleteTask deletes a task
func (s *TaskService) DeleteTask(id string) error {
	return s.repo.Delete(id)
}
