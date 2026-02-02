import 'package:tasks_flutter_web/domain/entities/task.dart';
import 'package:tasks_flutter_web/domain/repositories/task_repository.dart';

class ToggleTask {
  final TaskRepository repository;

  ToggleTask({required this.repository});

  Future<Task> call(String id) async {
    return await repository.toggleTaskCompletion(id);
  }
}
