import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:tasks_flutter_web/core/di/injection.dart';
import 'package:tasks_flutter_web/presentation/cubit/task_cubit.dart';
import 'package:tasks_flutter_web/presentation/pages/tasks_page.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  // Setup dependency injection (like your DI setup in main.dart)
  await setupDependencies();

  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Tasks App',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.blue),
        useMaterial3: true,
      ),
      home: BlocProvider(
        create: (_) => getIt<TaskCubit>(),
        child: const TasksPage(),
      ),
    );
  }
}
