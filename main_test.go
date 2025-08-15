package main_test

import (
	"context"
	"log"
	"testing"

	"github.com/1garo/manta/user"
	"github.com/1garo/manta/workflow"
)

func TestWorkflow(t *testing.T) {
	workflow := workflow.New()
	u := user.User{
		Name: "Alexandre",
		Age:  24,
	}
	workflow.Ctx = user.NewContext(workflow.Ctx, &u)

	workflow.AddTask("greet", func(ctx context.Context) error {
		if u, ok := user.FromContext(ctx); ok {
			log.Printf("grettings for %s which age is %d.\n", u.Name, u.Age)
		} else {
			log.Printf("user not found")
		}
		return nil
	})

	err := workflow.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
