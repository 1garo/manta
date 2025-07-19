package runner



import (
	"fmt"
)

// TaskFunc signature required for the task function to be executed
type TaskFunc func()

// Runner is the container for a specific set of tasks to be executed
//
// Tasks are isolated setup (right now, ignoring the order). If you want, a task can have dependencies.
// 
type Runner struct {
	pool map[string]TaskFunc
	deps map[string][]string
}

// New create a new runner
func New() *Runner {
	return &Runner{
		pool: make(map[string]TaskFunc),
		deps: make(map[string][]string),
	}
}

// AddTask add a task to the runner
func (r *Runner) AddTask(name string, fn TaskFunc) {
	r.pool[name] = fn
}

// AddDependency add a dependecy task to another task
func (r *Runner) AddDependency(name string, dependency string) {
	r.deps[name] = append(r.deps[name], dependency)
}

// AddDependencies add multiple dependecies task to another task
func (r *Runner) AddDependencies(name string, dependencies []string) {
	for _, dep := range dependencies {
		r.AddDependency(name, dep)
	}
}

// Run run the task attached function
func (r *Runner) Run(name string) error {
	if ok := r.hasDependencies(name); ok {
		return r.runWithDeps(name)
	}

	return r.run(name)
}

// hasDependencies check if the task has dependencies
func (r *Runner) hasDependencies(name string) bool {
	_, ok := r.deps[name]
	return ok
}

// runWithDeps run a task that has dependencies
func (r *Runner) runWithDeps(name string) error { 
	for _, dep := range r.deps[name] {
		if err := r.run(dep); err != nil {
			return err
		}
	}
	defer r.run(name)

	return nil
}

// run execute the function for the particular task
func (r *Runner) run(task string) error {
	fn, ok := r.pool[task]
	if !ok {
		return fmt.Errorf("task %s not found", task)
	}
	// TODO: how we want to capture the error that happen inside the executed function and expose this to the user?
	// We could consider changing the TaskFunc signature and force an error to be returned.
	// For logging across all of the tasks, we possibly could introduce a ctx that contains the logger.
	defer fn()

	return nil
}


