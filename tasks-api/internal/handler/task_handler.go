package handler

import (
	"encoding/json"
	"net/http"
	"tasks-api/internal/service"
)

// TaskHandler handles HTTP requests (similar to controllers in other frameworks)
type TaskHandler struct {
	service *service.TaskService
}

// NewTaskHandler creates a new task handler
func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

// CreateTaskRequest represents the request body for creating a task
type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// CreateTask handles POST /tasks
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	task, err := h.service.CreateTask(req.Title, req.Description)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.sendJSON(w, task, http.StatusCreated)
}

// GetAllTasks handles GET /tasks
func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.sendError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tasks, err := h.service.GetAllTasks()
	if err != nil {
		h.sendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.sendJSON(w, tasks, http.StatusOK)
}

// GetTaskByID handles GET /tasks/{id}
func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.sendError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		h.sendError(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusNotFound)
		return
	}

	h.sendJSON(w, task, http.StatusOK)
}

// ToggleTaskCompletion handles PUT /tasks/{id}/toggle
func (h *TaskHandler) ToggleTaskCompletion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		h.sendError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		h.sendError(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	task, err := h.service.ToggleTaskCompletion(id)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusNotFound)
		return
	}

	h.sendJSON(w, task, http.StatusOK)
}

// DeleteTask handles DELETE /tasks/{id}
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.sendError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		h.sendError(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteTask(id)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Helper functions
func (h *TaskHandler) sendJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func (h *TaskHandler) sendError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
