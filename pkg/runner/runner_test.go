package runner_test

import (
	"fmt"
	"testing"

	"github.com/1garo/manta/pkg/runner"
	"github.com/1garo/manta/pkg/task"
)

func TestRunner(t *testing.T) {
	taskName := "countNumber"
	taskFn := func () error {
		count := 0
		for _, n := range []int{1,2,3,4} {
			count += n
		}
		fmt.Println("count: ", count)
		return nil
	}
	tk := task.NewTask(taskName, taskFn)

	r := runner.New([]task.Task{tk})
	if _, ok := r.Execute(); !ok {
		t.Fatalf("expected %v, got %v", true, ok)
	}

}
