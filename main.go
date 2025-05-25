package main

import "fmt"

type TaskFunc func()
type Runner struct {
	pool map[string]TaskFunc
}

func New() *Runner {
	return &Runner{pool: make(map[string]TaskFunc)}
}

func (r *Runner) Task(name string, fn TaskFunc) {
	r.pool[name] = fn
}

func (r *Runner) Run(name string) error {
	fn, ok := r.pool[name]
	if !ok {
		return fmt.Errorf("task %q not found", name)
	}
	fn()
	return nil
}

func main() {
	fmt.Println("hellow world")
}
