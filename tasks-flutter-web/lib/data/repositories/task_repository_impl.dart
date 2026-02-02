import 'package:tasks_flutter_web/data/datasources/task_remote_datasource.dart';
import 'package:tasks_flutter_web/domain/entities/task.dart';
import 'package:tasks_flutter_web/domain/repositories/task_repository.dart';

/// Repository implementation - bridges data layer and domain layer
class TaskRepositoryImpl implements TaskRepository {
  final TaskRemoteDataSource remoteDataSource;

  TaskRepositoryImpl({required this.remoteDataSource});

  @override
  Future<List<Task>> getAllTasks() async {
    final taskModels = await remoteDataSource.getAllTasks();
    return taskModels.map((model) => model.toEntity()).toList();
  }

  @override
  Future<Task> createTask(String title, String description) async {
    final taskModel = await remoteDataSource.createTask(title, description);
    return taskModel.toEntity();
  }

  @override
  Future<Task> toggleTaskCompletion(String id) async {
    final taskModel = await remoteDataSource.toggleTaskCompletion(id);
    return taskModel.toEntity();
  }

  @override
  Future<void> deleteTask(String id) async {
    await remoteDataSource.deleteTask(id);
  }
}
