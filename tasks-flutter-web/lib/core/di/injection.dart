import 'package:dio/dio.dart';
import 'package:get_it/get_it.dart';
import 'package:tasks_flutter_web/data/datasources/task_remote_datasource.dart';
import 'package:tasks_flutter_web/data/repositories/task_repository_impl.dart';
import 'package:tasks_flutter_web/domain/repositories/task_repository.dart';
import 'package:tasks_flutter_web/domain/usecases/create_task.dart';
import 'package:tasks_flutter_web/domain/usecases/delete_task.dart';
import 'package:tasks_flutter_web/domain/usecases/get_all_tasks.dart';
import 'package:tasks_flutter_web/domain/usecases/toggle_task.dart';
import 'package:tasks_flutter_web/presentation/cubit/task_cubit.dart';

/// Dependency Injection setup (like your get_it setup)
final getIt = GetIt.instance;

Future<void> setupDependencies() async {
  // Dio instance (HTTP client)
  getIt.registerLazySingleton<Dio>(
    () => Dio(
      BaseOptions(
        baseUrl: 'http://localhost:8080', // Your Go backend URL
        connectTimeout: const Duration(seconds: 5),
        receiveTimeout: const Duration(seconds: 3),
      ),
    ),
  );

  // Data sources
  getIt.registerLazySingleton<TaskRemoteDataSource>(
    () => TaskRemoteDataSourceImpl(dio: getIt()),
  );

  // Repositories
  getIt.registerLazySingleton<TaskRepository>(
    () => TaskRepositoryImpl(remoteDataSource: getIt()),
  );

  // Use cases
  getIt.registerLazySingleton(() => GetAllTasks(repository: getIt()));
  getIt.registerLazySingleton(() => CreateTask(repository: getIt()));
  getIt.registerLazySingleton(() => ToggleTask(repository: getIt()));
  getIt.registerLazySingleton(() => DeleteTask(repository: getIt()));

  // Cubit
  getIt.registerFactory(
    () => TaskCubit(
      getAllTasks: getIt(),
      createTask: getIt(),
      toggleTask: getIt(),
      deleteTask: getIt(),
    ),
  );
}
