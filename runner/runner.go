package runner

import "github.com/1garo/manta/task"

// Runner represents the container of tasks.
type Runner struct {
	tasks map[string]*task.Task
}


// NewRunner create a new scheduler instance.
func NewRunner() *Runner {
	return &Runner{
		tasks: make(map[string]*task.Task, 0),
	}
}

// AddTask add an task to the runner.
func (s *Runner) AddTask(task *task.Task) error {
	s.tasks[task.Name] = task
	return nil
}

// Len get the length of the current scheduler.
func (s *Runner) Len() int {
	return len(s.tasks)
}

func (s *Runner) Execute() error {
	return nil
}

