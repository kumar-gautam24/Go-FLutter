package domain

import "time"

// Task represents a task entity (similar to your Flutter domain entities)
type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"isCompleted"`
	CreatedAt   time.Time `json:"createdAt"`
}

// NewTask creates a new task with generated ID
func NewTask(title, description string) *Task {
	return &Task{
		ID:          generateID(),
		Title:       title,
		Description: description,
		IsCompleted: false,
		CreatedAt:   time.Now(),
	}
}

// generateID generates a simple ID (in real apps, use UUID)
func generateID() string {
	return time.Now().Format("20060102150405")
}
