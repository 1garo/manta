package runner

import "github.com/1garo/manta/pkg/task"

type Runner struct {
	tasks []task.Task
}

func New(tasks []task.Task) *Runner{
	return &Runner{
		tasks,
	}
}

func (r *Runner) Execute() (map[string]error, bool){
	errs := make(map[string]error)
	for _, t := range r.tasks {
		if err := t.Run(); err != nil {
			errs[t.Name] = err
		}
	}

	if len(errs) > 0 {
		return errs, false
	}
	return errs, true
}
