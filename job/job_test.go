package job

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func handlePrintHelloWorld() error {
	fmt.Println("Hello world")
	return nil
}

var _ = Describe("Job", func() {
	Describe("Run a job to print hello world", func () {
		job := NewJob("PrintHelloWorld", handlePrintHelloWorld)
		Context("when input is valid", func() {
			It("should print the message", func() {
				err := job.Exec()
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
