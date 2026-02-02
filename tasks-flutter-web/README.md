# Tasks Flutter Web

A Flutter web app that demonstrates clean architecture while calling a Go backend API.

## Project Structure

```
lib/
├── core/di/                  # Dependency Injection (get_it)
│   └── injection.dart
├── domain/                   # Business logic layer (pure Dart)
│   ├── entities/
│   │   └── task.dart         # Domain entity
│   ├── repositories/
│   │   └── task_repository.dart  # Repository interface
│   └── usecases/             # Use cases (business actions)
│       ├── get_all_tasks.dart
│       ├── create_task.dart
│       ├── toggle_task.dart
│       └── delete_task.dart
├── data/                     # Data layer
│   ├── models/
│   │   └── task_model.dart   # JSON serialization
│   ├── datasources/
│   │   └── task_remote_datasource.dart  # API calls (Dio)
│   └── repositories/
│       └── task_repository_impl.dart  # Repository implementation
├── presentation/             # UI layer
│   ├── cubit/
│   │   ├── task_cubit.dart   # State management
│   │   └── task_state.dart   # States
│   └── pages/
│       └── tasks_page.dart   # Main UI
└── main.dart                 # Entry point
```

## Architecture Layers

### 1. Domain Layer (Business Logic)
**Pure Dart, no dependencies**

- **Entities**: Business objects
  ```dart
  class Task extends Equatable {
    final String id;
    final String title;
    // ...
  }
  ```

- **Repository Interfaces**: Abstract contracts
  ```dart
  abstract class TaskRepository {
    Future<List<Task>> getAllTasks();
    Future<Task> createTask(String title, String description);
  }
  ```

- **Use Cases**: Business actions
  ```dart
  class GetAllTasks {
    final TaskRepository repository;

    Future<List<Task>> call() async {
      return await repository.getAllTasks();
    }
  }
  ```

### 2. Data Layer
**Handles external data sources**

- **Models**: Add JSON serialization to entities
  ```dart
  class TaskModel extends Task {
    factory TaskModel.fromJson(Map<String, dynamic> json) {
      return TaskModel(
        id: json['id'],
        title: json['title'],
        // ...
      );
    }
  }
  ```

- **Data Sources**: API calls using Dio
  ```dart
  class TaskRemoteDataSourceImpl {
    final Dio dio;

    Future<List<TaskModel>> getAllTasks() async {
      final response = await dio.get('/tasks');
      return (response.data as List)
          .map((json) => TaskModel.fromJson(json))
          .toList();
    }
  }
  ```

- **Repository Implementation**: Bridges data and domain
  ```dart
  class TaskRepositoryImpl implements TaskRepository {
    final TaskRemoteDataSource remoteDataSource;

    @override
    Future<List<Task>> getAllTasks() async {
      final models = await remoteDataSource.getAllTasks();
      return models.map((m) => m.toEntity()).toList();
    }
  }
  ```

### 3. Presentation Layer
**UI and state management**

- **States**: Different UI states
  ```dart
  abstract class TaskState extends Equatable {}

  class TaskLoading extends TaskState {}
  class TaskLoaded extends TaskState {
    final List<Task> tasks;
  }
  class TaskError extends TaskState {
    final String message;
  }
  ```

- **Cubit**: Manages state
  ```dart
  class TaskCubit extends Cubit<TaskState> {
    final GetAllTasks getAllTasks;

    Future<void> loadTasks() async {
      emit(TaskLoading());
      final tasks = await getAllTasks();
      emit(TaskLoaded(tasks: tasks));
    }
  }
  ```

- **Pages**: UI widgets
  ```dart
  class TasksPage extends StatelessWidget {
    @override
    Widget build(BuildContext context) {
      return BlocBuilder<TaskCubit, TaskState>(
        builder: (context, state) {
          if (state is TaskLoading) {
            return CircularProgressIndicator();
          }
          // ...
        },
      );
    }
  }
  ```

## Data Flow

### Loading Tasks
```
1. UI triggers: context.read<TaskCubit>().loadTasks()
2. Cubit calls: GetAllTasks use case
3. Use case calls: TaskRepository.getAllTasks()
4. Repository calls: RemoteDataSource.getAllTasks()
5. DataSource makes: HTTP GET request to Go API
6. Response flows back through layers
7. Cubit emits: TaskLoaded state
8. UI rebuilds with new data
```

### Creating a Task
```
1. UI: User fills form and clicks "Add"
2. Cubit: addTask(title, description)
3. Use case: CreateTask
4. Repository: createTask()
5. DataSource: POST request to Go API
6. Response: New task from backend
7. Cubit: Reload all tasks
8. UI: Shows updated list
```

## Dependency Injection Setup

In [injection.dart](lib/core/di/injection.dart):

```dart
// 1. Register Dio
getIt.registerLazySingleton<Dio>(() => Dio(
  BaseOptions(baseUrl: 'http://localhost:8080'),
));

// 2. Register data sources
getIt.registerLazySingleton<TaskRemoteDataSource>(
  () => TaskRemoteDataSourceImpl(dio: getIt()),
);

// 3. Register repositories
getIt.registerLazySingleton<TaskRepository>(
  () => TaskRepositoryImpl(remoteDataSource: getIt()),
);

// 4. Register use cases
getIt.registerLazySingleton(() => GetAllTasks(repository: getIt()));

// 5. Register cubit (factory - new instance each time)
getIt.registerFactory(() => TaskCubit(
  getAllTasks: getIt(),
  createTask: getIt(),
  // ...
));
```

## Running the App

1. Make sure Go backend is running on `http://localhost:8080`

2. Get dependencies:
   ```bash
   flutter pub get
   ```

3. Run on Chrome:
   ```bash
   flutter run -d chrome
   ```

## Key Features

- **Clean Architecture** - Separation of concerns
- **State Management** - Using Cubit (part of bloc)
- **Dependency Injection** - Using get_it
- **HTTP Client** - Using Dio (like Retrofit)
- **Reactive UI** - BlocBuilder rebuilds on state changes

## Learning Points

### Clean Architecture Benefits
- **Testable**: Each layer can be tested independently
- **Maintainable**: Changes in one layer don't affect others
- **Scalable**: Easy to add new features
- **Understandable**: Clear separation of responsibilities

### Why So Many Files?
You might think "this is complex for a simple app!" but:
- In real apps, this structure prevents chaos
- Each file has ONE responsibility
- Easy to find and modify code
- Team members know where to add features

### Comparison to Your Usual Setup
You're familiar with this pattern! It's the same as your Flutter projects:
- **Domain entities** ✅ You use these
- **Repository pattern** ✅ You use this
- **Use cases** ✅ You use these
- **Cubit** ✅ You use this
- **Dio** ✅ You use this
- **get_it** ✅ You use this

The only difference: This calls a **Go backend** instead of a regular REST API!

## What's Different from Your Flutter Projects?

1. **No json_serializable** - Manual JSON parsing (you can add it if you want)
2. **Calling local Go API** - Instead of remote API
3. **Simpler models** - No complex nested objects yet

## Next Steps

1. **Add json_serializable** - Generate fromJson/toJson
2. **Add error handling** - Better error states
3. **Add loading indicators** - Per-task operations
4. **Add offline support** - Local caching
5. **Add authentication** - Login/logout flow

Enjoy building! 🎨
