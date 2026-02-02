import 'package:tasks_flutter_web/domain/entities/task.dart';
import 'package:tasks_flutter_web/domain/repositories/task_repository.dart';

/// Use case - encapsulates business logic (like your use cases)
class GetAllTasks {
  final TaskRepository repository;

  GetAllTasks({required this.repository});

  Future<List<Task>> call() async {
    return await repository.getAllTasks();
  }
}
