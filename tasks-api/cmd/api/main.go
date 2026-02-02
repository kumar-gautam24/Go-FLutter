package main

import (
	"fmt"
	"log"
	"net/http"
	"tasks-api/internal/handler"
	"tasks-api/internal/repository"
	"tasks-api/internal/service"
)

func main() {
	// Dependency Injection (like you do in Flutter with get_it)
	// 1. Create repository
	taskRepo := repository.NewInMemoryTaskRepository()

	// 2. Create service with repository
	taskService := service.NewTaskService(taskRepo)

	// 3. Create handler with service
	taskHandler := handler.NewTaskHandler(taskService)

	// Set up routes
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS for Flutter web
		enableCORS(&w)

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Route based on HTTP method
		switch r.Method {
		case http.MethodGet:
			taskHandler.GetAllTasks(w, r)
		case http.MethodPost:
			taskHandler.CreateTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/tasks/by-id", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		taskHandler.GetTaskByID(w, r)
	})

	mux.HandleFunc("/tasks/toggle", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		taskHandler.ToggleTaskCompletion(w, r)
	})

	mux.HandleFunc("/tasks/delete", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		taskHandler.DeleteTask(w, r)
	})

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy"}`))
	})

	// Start server
	port := "8080"
	fmt.Printf("🚀 Server starting on http://localhost:%s\n", port)
	fmt.Println("📝 Available endpoints:")
	fmt.Println("  GET    /tasks           - Get all tasks")
	fmt.Println("  POST   /tasks           - Create a task")
	fmt.Println("  GET    /tasks/by-id?id= - Get task by ID")
	fmt.Println("  PUT    /tasks/toggle?id= - Toggle task completion")
	fmt.Println("  DELETE /tasks/delete?id= - Delete task")
	fmt.Println("  GET    /health          - Health check")

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

// enableCORS adds CORS headers to allow Flutter web to call the API
func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
