import 'package:equatable/equatable.dart';

/// Domain entity - pure business object
class Task extends Equatable {
  final String id;
  final String title;
  final String description;
  final bool isCompleted;
  final DateTime createdAt;

  const Task({
    required this.id,
    required this.title,
    required this.description,
    required this.isCompleted,
    required this.createdAt,
  });

  @override
  List<Object?> get props => [id, title, description, isCompleted, createdAt];
}
