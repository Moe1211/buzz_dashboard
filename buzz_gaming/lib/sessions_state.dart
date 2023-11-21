// sessions_state.dart
part of 'sessions_bloc.dart';

abstract class SessionsState extends Equatable {
  const SessionsState();

  @override
  List<Object> get props => [];
}

class SessionsInitial extends SessionsState {}

class SessionStartedState extends SessionsState {}

class SessionsHistoryLoadedState extends SessionsState {
  final List<dynamic> history;

  const SessionsHistoryLoadedState({required this.history});

  @override
  List<Object> get props => [history];
}

class SessionErrorState extends SessionsState {
  final String message;

  const SessionErrorState({required this.message});

  @override
  List<Object> get props => [message];
}
