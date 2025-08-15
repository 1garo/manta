package workflow

import (
	"context"
	"log"
)

type Fn func (ctx context.Context) error
// Workflow contains one or more steps
type Workflow struct{
	pool map[string]Fn
	Ctx context.Context
}

// New returns a Workflow
func New() Workflow {
	return Workflow{
		pool: make(map[string]Fn, 0),
		Ctx: context.TODO(),
	}
}


func (w *Workflow) Execute() error {
	for name, task := range w.pool {
		log.Printf("Executing task: %s\n", name)
		return task(w.Ctx)
	}
	log.Println("finished executing")
	return nil
} 

func (w *Workflow) AddTask(name string, fn Fn) {
	w.pool[name] = fn
}
