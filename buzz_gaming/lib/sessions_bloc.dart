// sessions_bloc.dart
import 'dart:convert';

import 'package:equatable/equatable.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:http/http.dart' as http;

part 'sessions_event.dart';
part 'sessions_state.dart';

class SessionsBloc extends Bloc<SessionsEvent, SessionsState> {
  SessionsBloc() : super(SessionsInitial()) {
    on<StartSessionEvent>(startSessionEvent);
    on<ViewSessionsHistoryEvent>(viewSessions);
  }

  startSessionEvent(StartSessionEvent event, emit) async {
    try {
      // Make API request to start a session
      // You can use Dio, http package, or any other HTTP client
      final response =
          await http.post(Uri(scheme: "localhost:8888/user/startsession")
              // Provide necessary headers, body, etc.
              );

      if (response.statusCode == 200) {
        // Session started successfully
        return SessionStartedState();
      } else {
        // Handle error
        return const SessionErrorState(message: 'Failed to start session');
      }
    } catch (e) {
      // Handle network error
      return const SessionErrorState(message: 'Network error');
    }
  }

  viewSessions(ViewSessionsHistoryEvent event, emit) async {
    try {
      // Make API request to view sessions history
      final response = await http
          .post(Uri(scheme: "localhost:8888/user/viewsessionshistory"));

      if (response.statusCode == 200) {
        // Parse the response and handle the sessions history
        final List<dynamic> sessionsHistory = json.decode(response.body);
        return SessionsHistoryLoadedState(history: sessionsHistory);
      } else {
        // Handle error
        return const SessionErrorState(
            message: 'Failed to view sessions history');
      }
    } catch (e) {
      // Handle network error
      return const SessionErrorState(message: 'Network error');
    }
  }
}
