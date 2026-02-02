import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:tasks_flutter_web/domain/usecases/create_task.dart';
import 'package:tasks_flutter_web/domain/usecases/delete_task.dart';
import 'package:tasks_flutter_web/domain/usecases/get_all_tasks.dart';
import 'package:tasks_flutter_web/domain/usecases/toggle_task.dart';
import 'package:tasks_flutter_web/presentation/cubit/task_state.dart';

/// Cubit for managing task state (like your Cubits)
class TaskCubit extends Cubit<TaskState> {
  final GetAllTasks getAllTasks;
  final CreateTask createTask;
  final ToggleTask toggleTask;
  final DeleteTask deleteTask;

  TaskCubit({
    required this.getAllTasks,
    required this.createTask,
    required this.toggleTask,
    required this.deleteTask,
  }) : super(TaskInitial());

  Future<void> loadTasks() async {
    try {
      emit(TaskLoading());
      final tasks = await getAllTasks();
      emit(TaskLoaded(tasks: tasks));
    } catch (e) {
      emit(TaskError(message: e.toString()));
    }
  }

  Future<void> addTask(String title, String description) async {
    try {
      await createTask(title, description);
      await loadTasks(); // Reload tasks after creating
    } catch (e) {
      emit(TaskError(message: e.toString()));
    }
  }

  Future<void> toggleTaskCompletion(String id) async {
    try {
      await toggleTask(id);
      await loadTasks(); // Reload tasks after toggling
    } catch (e) {
      emit(TaskError(message: e.toString()));
    }
  }

  Future<void> removeTask(String id) async {
    try {
      await deleteTask(id);
      await loadTasks(); // Reload tasks after deleting
    } catch (e) {
      emit(TaskError(message: e.toString()));
    }
  }
}
