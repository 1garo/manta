package task_test

import (
	"testing"
	"fmt"
	"github.com/1garo/manta/pkg/task"
)

func TestRunTask(t *testing.T) {
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
	err := tk.Run()

	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	tk = task.NewTask("failingTask", func() error {
		return fmt.Errorf("this task is going to fail")
	})

	err = tk.Run()
	if err == nil {
		t.Fatalf("expected error, but got %v", err)
	}

}
