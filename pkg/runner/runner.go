package runner

import "log"

type Runner struct {}

type Task struct {
	name string
	fn func() error
	failed bool
}

func New() *Runner{
	return &Runner{}
}

func (r *Runner) Task(name string, fn func() error) Task {
	return Task{
		name: name,
		fn: fn,
		failed: false,
	}
}

func (t *Task) Run() error {
	log.Printf("running the following task: %s\n", t.name)
	if err := t.fn(); err != nil {
		t.failed = true
		return nil
	}

	return nil
}

func (t *Task) Failed() bool {
	return t.failed
}
