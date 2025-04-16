# Manta
Your favorite job scheduler

## Job Scheduler with Plugin Adapters
Goal: Build a scheduler that runs different job types (e.g., HTTP calls, shell commands) concurrently.

Concurrency Focus:
- Manage job execution with goroutines and track progress using channels.
- Use `context.Context` for cancellations/timeouts.

Adapter Pattern:
- Define a Job interface and adapters for specific job types:
```go
type Job interface {
  Execute() error
}

type HTTPJob struct { URL string } // Implements Execute()
type ShellJob struct { Command string } // Implements Execute()
```
