package main

import (
	"fmt"
	"log"
	"net/http"
	"tasks-api/internal/handler"
	"tasks-api/internal/repository"
	"tasks-api/internal/router"
	"tasks-api/internal/service"
)

const port = "8080"

func main() {
	taskRepo := repository.NewInMemoryTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	mux := router.Setup(taskHandler)

	logServerStart()

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

func logServerStart() {
	fmt.Printf("Server starting on http://localhost:%s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET    /tasks           - Get all tasks")
	fmt.Println("  POST   /tasks           - Create a task")
	fmt.Println("  GET    /tasks/by-id?id= - Get task by ID")
	fmt.Println("  PUT    /tasks/toggle?id= - Toggle task completion")
	fmt.Println("  DELETE /tasks/delete?id= - Delete task")
	fmt.Println("  GET    /health          - Health check")
}
