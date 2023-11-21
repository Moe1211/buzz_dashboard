// main.dart
import 'dart:io';

import 'package:device_preview/device_preview.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import 'sessions_bloc.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return DevicePreview(
        enabled: Platform.isLinux || Platform.isMacOS,
        builder: (context) {
          return MultiBlocProvider(
            providers: [
              BlocProvider<SessionsBloc>(create: (context) => SessionsBloc())
            ],
            child: MaterialApp(
              title: 'Flutter BLoC Example',
              useInheritedMediaQuery: true,
              theme: ThemeData(
                primarySwatch: Colors.blue,
              ),
              home: const MyHomePage(),
            ),
          );
        });
  }
}

class MyHomePage extends StatelessWidget {
  const MyHomePage({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => SessionsBloc(),
      child: Scaffold(
        appBar: AppBar(
          title: const Text('Flutter BLoC Example'),
        ),
        body: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              ElevatedButton(
                onPressed: () {
                  // Dispatch the StartSessionEvent
                  BlocProvider.of<SessionsBloc>(context)
                      .add(StartSessionEvent());
                },
                child: const Text('Start Session'),
              ),
              const SizedBox(height: 20),
              ElevatedButton(
                onPressed: () {
                  // Dispatch the ViewSessionsHistoryEvent
                  BlocProvider.of<SessionsBloc>(context)
                      .add(ViewSessionsHistoryEvent());
                },
                child: const Text('View Sessions History'),
              ),
              const SizedBox(height: 20),
              BlocBuilder<SessionsBloc, SessionsState>(
                builder: (context, state) {
                  if (state is SessionsHistoryLoadedState) {
                    // Display sessions history
                    return Column(
                      children: [
                        const Text('Sessions History:'),
                        for (var session in state.history)
                          ListTile(
                            title: Text(session['start_time']),
                            subtitle: Text(session['end_time'] ?? 'Ongoing'),
                          ),
                      ],
                    );
                  } else if (state is SessionErrorState) {
                    // Display error message
                    return Text('Error: ${state.message}');
                  } else {
                    // Display loading or initial state
                    return const CircularProgressIndicator();
                  }
                },
              ),
            ],
          ),
        ),
      ),
    );
  }
}
