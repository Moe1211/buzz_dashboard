// sessions_event.dart
part of 'sessions_bloc.dart';

abstract class SessionsEvent extends Equatable {
  const SessionsEvent();

  @override
  List<Object> get props => [];
}

class StartSessionEvent extends SessionsEvent {}

class ViewSessionsHistoryEvent extends SessionsEvent {}
