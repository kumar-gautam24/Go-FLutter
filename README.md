# Go + Flutter Learning Project

This folder contains two projects to help you learn Go backend development:

1. **tasks-api** - A Go REST API backend
2. **tasks-flutter-web** - A Flutter web app that consumes the Go API

## Architecture Overview

Both projects follow **Clean Architecture** principles that you're familiar with from Flutter:

### Go Backend Structure
```
tasks-api/
├── cmd/api/              # Entry point (like main.dart)
│   └── main.go           # Server setup & dependency injection
├── internal/
│   ├── domain/           # Business entities (like domain layer in Flutter)
│   │   └── task.go       # Task entity
│   ├── repository/       # Data layer (like repositories in Flutter)
│   │   └── task_repository.go  # In-memory storage
│   ├── service/          # Business logic (like use cases in Flutter)
│   │   └── task_service.go
│   └── handler/          # HTTP handlers (like controllers)
│       └── task_handler.go
└── go.mod                # Dependencies (like pubspec.yaml)
```

### Flutter Web Structure
```
tasks-flutter-web/
├── lib/
│   ├── core/di/          # Dependency injection (get_it)
│   ├── domain/           # Business logic layer
│   │   ├── entities/     # Domain entities
│   │   ├── repositories/ # Repository interfaces
│   │   └── usecases/     # Use cases
│   ├── data/             # Data layer
│   │   ├── models/       # Data models with JSON serialization
│   │   ├── datasources/  # API calls (using Dio)
│   │   └── repositories/ # Repository implementations
│   └── presentation/     # UI layer
│       ├── cubit/        # State management (Cubit)
│       └── pages/        # UI pages
└── pubspec.yaml
```

## How to Run

### Step 1: Run the Go Backend

1. Open a terminal and navigate to the Go project:
   ```bash
   cd tasks-api
   ```

2. Run the server:
   ```bash
   go run cmd/api/main.go
   ```

3. You should see:
   ```
   🚀 Server starting on http://localhost:8080
   📝 Available endpoints:
     GET    /tasks           - Get all tasks
     POST   /tasks           - Create a task
     GET    /tasks/by-id?id= - Get task by ID
     PUT    /tasks/toggle?id= - Toggle task completion
     DELETE /tasks/delete?id= - Delete task
     GET    /health          - Health check
   ```

### Step 2: Run the Flutter Web App

1. Open a **new** terminal and navigate to the Flutter project:
   ```bash
   cd tasks-flutter-web
   ```

2. Get dependencies:
   ```bash
   flutter pub get
   ```

3. Run the app:
   ```bash
   flutter run -d chrome
   ```

4. The app will open in Chrome and connect to your Go backend!

## Learning Guide

### For Go (Coming from Dart/C++)

#### Key Differences from Dart:

1. **No classes** - Go uses structs and functions
   ```go
   // Dart: class Task { ... }
   // Go:
   type Task struct {
       ID string
       Title string
   }
   ```

2. **Interfaces are implicit** - No need to declare implementation
   ```go
   // Just implement the methods, no "implements" keyword needed
   type TaskRepository interface {
       Create(task *Task) error
   }
   ```

3. **Error handling** - Functions return errors explicitly
   ```go
   // No try-catch, errors are returned values
   task, err := service.GetTask(id)
   if err != nil {
       // handle error
   }
   ```

4. **Pointers** - Similar to C++
   ```go
   // * means pointer, & means address
   func Update(task *Task) { } // accepts pointer
   Update(&myTask)             // pass address
   ```

5. **No null** - Use `nil` instead
   ```go
   var task *Task = nil
   ```

#### Go Concepts to Learn:

- **Structs** - Like classes without methods (add methods separately)
- **Interfaces** - Define behavior, implemented implicitly
- **Goroutines** - Lightweight threads (not used in this simple example)
- **Channels** - For communication between goroutines (not used here)
- **Defer** - Runs code at end of function (like try-finally)
- **Mutex** - For thread-safe operations (used in repository)

### Project Flow

1. **Request comes in** → handler (HTTP layer)
2. **Handler calls** → service (business logic)
3. **Service calls** → repository (data layer)
4. **Repository** → stores/retrieves data

This is the same flow you use in Flutter with clean architecture!

## API Endpoints

Test the API using your browser or tools like Postman:

- `GET http://localhost:8080/tasks` - Get all tasks
- `POST http://localhost:8080/tasks` - Create task (send JSON body)
- `GET http://localhost:8080/tasks/by-id?id=xxx` - Get one task
- `PUT http://localhost:8080/tasks/toggle?id=xxx` - Toggle completion
- `DELETE http://localhost:8080/tasks/delete?id=xxx` - Delete task

## Next Steps

1. **Understand the flow**: Follow a request from Flutter → Go → back to Flutter
2. **Modify the code**: Add a "priority" field to tasks
3. **Learn Go basics**: Variables, functions, structs, interfaces
4. **Explore Go stdlib**: `net/http`, `encoding/json`

## Common Go Commands

```bash
go run main.go          # Run a Go file
go build               # Compile to executable
go mod init <name>     # Initialize a new module
go mod tidy            # Clean up dependencies
go fmt ./...           # Format all code
```

## Resources

- [Go Tour](https://go.dev/tour/) - Interactive Go tutorial
- [Go by Example](https://gobyexample.com/) - Learn by examples
- [Effective Go](https://go.dev/doc/effective_go) - Best practices

Happy learning! 🚀
