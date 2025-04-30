package job

// Fn represents an executable function that returns an error
type Fn = func() error

// Job represents a named executable task
type Job struct {
	Name string // Human-readable identifier for the job
	exec Fn    // Function to execute when running the job
	// TODO: what if we have something similar to crontab schedule?
}

// NewJob creates a new Job instance
func NewJob(name string, f Fn) *Job {
	return &Job{
		Name: name,
		exec: f,
	}
}

// Execute runs the job's function.
func (j *Job) Execute() error {
	return j.exec()
}
