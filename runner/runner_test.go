package runner_test

import (
	"io"
	"log/slog"
	"net/http"
	"testing"

	"github.com/1garo/manta/runner"
)

func TestFullValidationAndAPICallTask(t *testing.T) {
	r := runner.New()

	input := "hello"

	r.AddTask("validate-not-empty", func() {
		slog.Info("function executed", "task name", "validate-not-empty")
		if input == "" {
			t.Fatal("input is empty")
		}
	})

	r.AddTask("validate-length", func() {
		slog.Info("function executed", "task name", "validate-length")
		if len(input) > 10 {
			t.Fatal("input too long")
		}
	})

	r.AddTask("call-api", func() {
		resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/-1")
		if err != nil {
			slog.Error("failed to call the API", "error", err)
			t.Fatalf("API call failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {

			b, _ := io.ReadAll(resp.Body)
			slog.Error("failed to call the API", "status code", resp.StatusCode, "body", string(b))
			t.Fatalf("unexpected status code: %d", resp.StatusCode)
		}
		slog.Info("successfully called the API")
	})

	r.AddDependencies("call-api", []string{"validate-not-empty", "validate-length"})

	if err := r.Run("call-api"); err != nil {
		t.Fatalf("run failed: %v", err)
	}
}
