import 'package:dio/dio.dart';
import 'package:tasks_flutter_web/data/models/task_model.dart';

/// Remote data source - handles API calls (like your Retrofit services)
abstract class TaskRemoteDataSource {
  Future<List<TaskModel>> getAllTasks();
  Future<TaskModel> createTask(String title, String description);
  Future<TaskModel> toggleTaskCompletion(String id);
  Future<void> deleteTask(String id);
}

class TaskRemoteDataSourceImpl implements TaskRemoteDataSource {
  final Dio dio;

  TaskRemoteDataSourceImpl({required this.dio});

  @override
  Future<List<TaskModel>> getAllTasks() async {
    try {
      final response = await dio.get('/tasks');

      final List<dynamic> data = response.data as List<dynamic>;
      return data.map((json) => TaskModel.fromJson(json)).toList();
    } on DioException catch (e) {
      throw Exception('Failed to fetch tasks: ${e.message}');
    }
  }

  @override
  Future<TaskModel> createTask(String title, String description) async {
    try {
      final response = await dio.post(
        '/tasks',
        data: {
          'title': title,
          'description': description,
        },
      );

      return TaskModel.fromJson(response.data);
    } on DioException catch (e) {
      throw Exception('Failed to create task: ${e.message}');
    }
  }

  @override
  Future<TaskModel> toggleTaskCompletion(String id) async {
    try {
      final response = await dio.put('/tasks/toggle?id=$id');
      return TaskModel.fromJson(response.data);
    } on DioException catch (e) {
      throw Exception('Failed to toggle task: ${e.message}');
    }
  }

  @override
  Future<void> deleteTask(String id) async {
    try {
      await dio.delete('/tasks/delete?id=$id');
    } on DioException catch (e) {
      throw Exception('Failed to delete task: ${e.message}');
    }
  }
}
