package scheduler

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/1garo/manta/job"
)


var _ = Describe("Scheduler", func() {
	var sc *Scheduler
	var j *job.Job
	var jobName = "Hello world"

	BeforeEach(func() {
		sc = NewScheduler()
		j = job.NewJob(jobName, func () error {
			fmt.Println("hello world")
			return nil
		})
	})

	Describe("Creating a new scheduler", func () {
		Context("when scheduler is created correctly", func() {
			It("scheduler job lenght should be zero", func() {
				count := sc.Len()
				Expect(count).To(Equal(0))
			})
		})
		Context("when a new job is added", func() {
			It("scheduler job lenght should be one", func() {
				sc.AddJob(j)
				count := sc.Len()
				Expect(count).To(Equal(1))
			})
		})
	})

	Describe("Getting a job from a scheduler", func () {
		Context("when checking if a job exists", func() {
			It("should not found job", func() {
				job, ok := sc.Get("invalidName")
				Expect(ok).To(Equal(false))
				Expect(job).To(BeNil())
			})

			It("should found job named HelloWorld", func() {
				sc.AddJob(j)
				prevLen := sc.Len()

				job, ok := sc.Get(jobName)
				Expect(ok).To(Equal(true))
				Expect(job).ToNot(BeNil())

				currLen := sc.Len()
				Expect(currLen).To(Equal(prevLen - 1))
			})
		})
	})
})
