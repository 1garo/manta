package runner_test

import (
	"testing"
	"fmt"
	"github.com/1garo/manta/pkg/runner"
)

func TestRunSingleTask(t *testing.T) {
	r := runner.New()

	taskName := "countNumber"
	task := r.Task(taskName, func () error {
		count := 0
		for _, n := range []int{1,2,3,4} {
			count += n
		}
		fmt.Println("count: ", count)
		return nil
	})

	err := task.Run()

	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	if task.Failed() {
		t.Fatalf("expected '%s' task to be executed", taskName)
	}


}
