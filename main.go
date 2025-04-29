package main

import (
	"log"

	"github.com/1garo/manta/job"
	"github.com/1garo/manta/scheduler"
)

func main() {
	sc := scheduler.NewScheduler()

    log.Printf("queue length: %d", sc.Len())

	jobName := "HelloWorld"
	job := job.NewJob(jobName, func() error {
        log.Println("Hello World!")
        return nil
    })

	sc.AddJob(job)
    log.Printf("queue length before getting job: %d", sc.Len())
    j, ok := sc.Get(jobName)
    log.Printf("queue length after getting job: %d", sc.Len())
    if !ok {
        log.Fatal("job not found")
    }
	err := j.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
