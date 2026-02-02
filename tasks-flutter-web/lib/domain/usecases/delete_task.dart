import 'package:tasks_flutter_web/domain/repositories/task_repository.dart';

class DeleteTask {
  final TaskRepository repository;

  DeleteTask({required this.repository});

  Future<void> call(String id) async {
    await repository.deleteTask(id);
  }
}
