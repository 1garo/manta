package main_test

import (
	"net/http"
	"testing"
)

func TestFullValidationAndAPICallTask(t *testing.T) {
	r := New()

	input := "hello"

	r.Task("validate-not-empty", func() {
		if input == "" {
			t.Fatal("input is empty")
		}
	})

	r.Task("validate-length", func() {
		if len(input) > 10 {
			t.Fatal("input too long")
		}
	})

	r.TaskWithDeps("call-api", []string{"validate-not-empty", "validate-length"}, func() {
		resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
		if err != nil {
			t.Fatalf("API call failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			t.Fatalf("unexpected status code: %d", resp.StatusCode)
		}
	})

	if err := r.Run("call-api"); err != nil {
		t.Fatalf("run failed: %v", err)
	}
}
