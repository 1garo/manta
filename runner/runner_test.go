package runner

import (
	"fmt"

	"github.com/1garo/manta/task"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Runner", func() {
	var sc *Runner
	var j *task.Task

	var taskName = "hello"
	var successFn = func() error {
		fmt.Println("hello world")
		return nil
	}
	var failedFn = func() error {
		return fmt.Errorf("this function failed")
	}


	BeforeEach(func() {
		sc = NewRunner()
		j = task.NewTask(taskName)
	})

	Describe("Creating a new runner", func() {
		Context("when runner is created correctly", func() {
			It("runner lenght should be zero", func() {
				count := sc.Len()
				Expect(count).To(Equal(0))
			})
		})
		Context("when a new task is created", func() {
			It("runner should be one", func() {
				sc.AddTask(j)
				Expect(sc.Len()).To(Equal(1))
			})
		})

		Context("when a step that succeed is added to the task", func() {
			It("task size should be one", func() {
				s := task.NewStep("print hello world", successFn)
				j.AddStep(s)
				Expect(j.Size()).To(Equal(1))
			})

			It("task should be runned and steps executed", func() {
				s := task.NewStep("print hello world", successFn)
				j.AddStep(s)
				Expect(j.Size()).To(Equal(1))
				err := j.Run()
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when a step that fail is added to the task", func() {
			It("task size should be one", func() {
				s := task.NewStep("fail step", failedFn)
				j.AddStep(s)
				Expect(j.Size()).To(Equal(1))
			})

			It("task should be runned and steps execute but fail", func() {
				s := task.NewStep("fail step", failedFn)
				j.AddStep(s)
				Expect(j.Size()).To(Equal(1))
				err := j.Run()
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
