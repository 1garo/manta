package task

import "log"

type Fn func() error
type Task struct {
	Name   string
	fn    Fn 
	failed bool
}

func NewTask(name string, fn Fn) Task {
	return Task{
		Name: name,
		fn: fn,
		failed: false,
	}
}

func (t *Task) Run() error {
	log.Printf("running the following task: %s\n", t.Name)
	return t.fn()
}

// TODO: The following could be useful for the idea of having the multiple tasks
// (chatgpt helped)
//func (t *task) RunWithDeps() error {}
