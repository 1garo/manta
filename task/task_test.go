package task

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func successFn() error {
	fmt.Println("Hello world")
	return nil
}

func failedFn() error {
	return fmt.Errorf("this function failed")
}

var _ = Describe("Task", func() {
	var t *Task
	BeforeEach(func () {
		t = NewTask("TestTask")
	})

	Describe("Execute the steps of a task", func () {
		Context("when you execute the steps", func() {
			It("steps should succeed", func() {
				s := NewStep("success step", successFn)
				t.AddStep(s)
				Expect(t.Size()).To(Equal(1))
				err := t.Run()
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("steps should fail", func() {
				s := NewStep("fail step", failedFn)
				t.AddStep(s)
				Expect(t.Size()).To(Equal(1))
				err := t.Run()
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})
