package runner

import (
	"fmt"

	"github.com/1garo/manta/job"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Runner", func() {
	var sc *Runner
	var j *job.Job
	var jobName = "Hello world"

	var jobFn = func() error {
		fmt.Println("hello world")
		return nil
	}

	BeforeEach(func() {
		sc = NewRunner()
		j = job.NewJob(jobName, jobFn)
	})

	Describe("Creating a new runner", func() {
		Context("when runner is created correctly", func() {
			It("runner job lenght should be zero", func() {
				count := sc.Len()
				Expect(count).To(Equal(0))
			})
		})
		Context("when a new job is added", func() {
			It("runner job lenght should be one", func() {
				sc.AddJob(j)
				count := sc.Len()
				Expect(count).To(Equal(1))
			})
		})
	})

	Describe("Getting a job from a runner", func() {
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

	Describe("Bootstraping the runner", func() {
		Context("Calling bootstrap to add the jobs to runner", func() {
			It("Should create all the jobs correctly", func() {
				jobs := make(map[string]job.Fn, 0)
				jobs["firstJob"] = jobFn
				jobs["secondJob"] = jobFn
				Expect(sc.Len()).To(Equal(0))
				sc.Bootstrap(jobs)
				Expect(sc.Len()).To(Equal(2))
			})

		})
	})
})
