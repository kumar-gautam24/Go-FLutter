import 'package:tasks_flutter_web/domain/entities/task.dart';

/// Repository interface (abstract class - like your domain repositories)
abstract class TaskRepository {
  Future<List<Task>> getAllTasks();
  Future<Task> createTask(String title, String description);
  Future<Task> toggleTaskCompletion(String id);
  Future<void> deleteTask(String id);
}
