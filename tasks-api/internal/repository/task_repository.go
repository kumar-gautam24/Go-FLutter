package repository

import (
	"errors"
	"sync"
	"tasks-api/internal/domain"
)

type TaskRepository interface {
	Create(task *domain.Task) error
	GetAll() ([]*domain.Task, error)
	GetByID(id string) (*domain.Task, error)
	Update(task *domain.Task) error
	Delete(id string) error
}

type InMemoryTaskRepository struct {
	tasks map[string]*domain.Task
	mu    sync.RWMutex
}

func NewInMemoryTaskRepository() TaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]*domain.Task),
	}
}

func (r *InMemoryTaskRepository) Create(task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; exists {
		return errors.New("task already exists")
	}

	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryTaskRepository) GetAll() ([]*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	tasks := make([]*domain.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *InMemoryTaskRepository) GetByID(id string) (*domain.Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}

	return task, nil
}

func (r *InMemoryTaskRepository) Update(task *domain.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[task.ID]; !exists {
		return errors.New("task not found")
	}

	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryTaskRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return errors.New("task not found")
	}

	delete(r.tasks, id)
	return nil
}
