package task

import (
	"fmt"
	"log"
	"os"
)

// Fn represents an executable function that returns an error
type Fn = func() error

type Step struct {
	// Human-readable identifier for the task within a task
	Description string
	// Function to execute
	exec Fn
}

func NewStep(description string, f Fn) *Step {
	return &Step{
		Description: description,
		exec: f,
	}
}

// Task represents a executable task with one or more steps
type Task struct {
	// Human-readable identifier for the task
	Name string 
	// Steps to be executed
	steps []*Step
	logger *log.Logger
}

func NewTask(name string) *Task {
	return &Task{
		Name: name,
		logger: log.New(os.Stdout, "task: ", log.LstdFlags),
	}
}

func (t *Task) AddStep(s *Step) {
	// TODO: we can consider at some point have position here
	// so the user can determine where they want to put this new step
	t.steps = append(t.steps, s)
}

func (t *Task) Size() int {
	return len(t.steps)
}

// Run executes the t.steps[].exec, the order is guaranteed by the
// order you call task.AddStep
func (t *Task) Run() error {
	for _, s := range t.steps {
		t.logger.Println(fmt.Sprintf("trying to execute step: %s", s.Description))
		if err := s.exec(); err != nil {
			t.logger.Println(fmt.Sprintf("failed to execute step %s: %v", s.Description, err))
			return err
		}
		t.logger.Println("finished executing step")
	}
	return nil
}

// Execute runs the step function
func (s *Step) Execute() error {
	return s.exec()
}
