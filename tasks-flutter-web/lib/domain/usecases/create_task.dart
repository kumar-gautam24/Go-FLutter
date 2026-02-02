import 'package:tasks_flutter_web/domain/entities/task.dart';
import 'package:tasks_flutter_web/domain/repositories/task_repository.dart';

class CreateTask {
  final TaskRepository repository;

  CreateTask({required this.repository});

  Future<Task> call(String title, String description) async {
    return await repository.createTask(title, description);
  }
}
