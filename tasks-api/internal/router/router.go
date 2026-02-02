package router

import (
	"net/http"
	"tasks-api/internal/handler"
	"tasks-api/internal/middleware"
)

func Setup(taskHandler *handler.TaskHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			taskHandler.GetAllTasks(w, r)
		case http.MethodPost:
			taskHandler.CreateTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	mux.HandleFunc("/tasks/by-id", middleware.CORS(taskHandler.GetTaskByID))
	mux.HandleFunc("/tasks/toggle", middleware.CORS(taskHandler.ToggleTaskCompletion))
	mux.HandleFunc("/tasks/delete", middleware.CORS(taskHandler.DeleteTask))

	mux.HandleFunc("/health", middleware.CORS(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy"}`))
	}))

	return mux
}
