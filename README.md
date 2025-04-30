# Manta
Your favorite job runner.

## Job Runner
Goal: Build a service that runs jobs instantaneously and concurrently.

### Example

```go
package main

import (
	"log"

	"github.com/1garo/manta/job"
	"github.com/1garo/manta/runner"
)

func main() {
	sc := runner.NewRunner()

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
```

### Installation

- Run `go mod init github.com/<your-user>/<your-repo>`.
- Add the example content into `main.go` file.
- Run `go mod tidy`.
