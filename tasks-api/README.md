# Tasks API - Go Backend

A simple REST API built with Go to help you learn backend development.

## Project Structure

```
tasks-api/
├── cmd/api/main.go                      # Entry point, server setup
├── internal/
│   ├── domain/task.go                   # Task entity (business object)
│   ├── repository/task_repository.go    # Data access layer
│   ├── service/task_service.go          # Business logic
│   └── handler/task_handler.go          # HTTP request handlers
└── go.mod                               # Dependencies
```

## Clean Architecture Layers

### 1. Domain (`internal/domain/`)
- **Pure business entities**
- No dependencies on other layers
- Like your Flutter domain entities

```go
type Task struct {
    ID          string
    Title       string
    Description string
    IsCompleted bool
    CreatedAt   time.Time
}
```

### 2. Repository (`internal/repository/`)
- **Interface + Implementation**
- Handles data storage/retrieval
- In-memory storage for simplicity

```go
type TaskRepository interface {
    Create(task *Task) error
    GetAll() ([]*Task, error)
    // ...
}
```

### 3. Service (`internal/service/`)
- **Business logic / Use cases**
- Validates data
- Calls repository

```go
func (s *TaskService) CreateTask(title, description string) (*Task, error) {
    // Business validation
    if title == "" {
        return nil, errors.New("title cannot be empty")
    }
    // Create and store
    task := domain.NewTask(title, description)
    return task, s.repo.Create(task)
}
```

### 4. Handler (`internal/handler/`)
- **HTTP layer**
- Parses requests
- Calls service
- Returns responses

```go
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
    // Parse JSON body
    var req CreateTaskRequest
    json.NewDecoder(r.Body).Decode(&req)

    // Call service
    task, err := h.service.CreateTask(req.Title, req.Description)

    // Return response
    h.sendJSON(w, task, http.StatusCreated)
}
```

## How It Works

1. **Server starts** in `main.go`
2. **Dependencies injected** (repository → service → handler)
3. **Routes registered** (`/tasks`, `/tasks/by-id`, etc.)
4. **Requests handled**:
   - Handler receives HTTP request
   - Handler calls service
   - Service validates and calls repository
   - Repository stores/retrieves data
   - Response sent back through the layers

## Key Go Concepts Used

### Structs (like classes)
```go
type Task struct {
    ID    string
    Title string
}
```

### Methods (attached to structs)
```go
func (s *TaskService) CreateTask(title string) (*Task, error) {
    // s is the receiver (like 'this' in Dart)
}
```

### Interfaces (implicit implementation)
```go
type TaskRepository interface {
    Create(task *Task) error
}

// Any struct with Create method implements TaskRepository
// No explicit "implements" keyword needed
```

### Error Handling
```go
task, err := service.GetTask(id)
if err != nil {
    return err  // Handle error
}
// Use task
```

### Pointers
```go
func Update(task *Task) {  // accepts pointer
    task.Title = "New"     // modifies original
}

myTask := &Task{}          // create pointer
Update(myTask)             // pass pointer
```

### JSON Tags
```go
type Task struct {
    ID string `json:"id"`  // Maps to JSON field "id"
}
```

## Running the Server

```bash
# From tasks-api directory
go run cmd/api/main.go
```

Server starts on `http://localhost:8080`

## Testing Endpoints

### Get all tasks
```bash
curl http://localhost:8080/tasks
```

### Create a task
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go","description":"Study clean architecture"}'
```

### Toggle completion
```bash
curl -X PUT "http://localhost:8080/tasks/toggle?id=YOUR_TASK_ID"
```

### Delete task
```bash
curl -X DELETE "http://localhost:8080/tasks/delete?id=YOUR_TASK_ID"
```

## Learning Tips

1. **Start with `main.go`** - See how everything connects
2. **Follow a request** - Trace from handler → service → repository
3. **Understand interfaces** - They're different from Dart's interfaces
4. **Learn error handling** - No exceptions, errors are values
5. **Practice pointers** - Similar to C++, but safer

## What's Next?

After understanding this code:
- Add a database (PostgreSQL, MongoDB)
- Add authentication (JWT)
- Add validation (go-validator)
- Add tests (testing package)
- Learn goroutines (concurrency)

Happy coding! 🎯
