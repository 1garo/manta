package main

import (
	"log"

	"github.com/1garo/manta/task"
	"github.com/1garo/manta/runner"
)

func main() {
	sc := runner.NewRunner()

	log.Printf("runner size: %d", sc.Len())

	jobName := "HelloWorld"
	job := task.NewTask(jobName, func() error {
		log.Println("Hello World!")
		return nil
	})

	sc.AddTask(job)
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
