package runner

import "github.com/1garo/manta/job"

// Runner represents the container for the jobs.
type Runner struct {
	jobs map[string]*job.Job
}


// NewRunner create a new scheduler instance.
func NewRunner() *Runner {
	return &Runner{
		jobs: make(map[string]*job.Job, 0),
	}
}

// AddJob add a job to the current scheduler.
func (s *Runner) AddJob(job *job.Job) error {
	s.jobs[job.Name] = job
	return nil
}

// Len get the length of the current scheduler.
func (s *Runner) Len() int {
	return len(s.jobs)
}

// Get will if the jobs exists, pop from the queue.
//
// If not, it will returns nil and false.
func (s *Runner) Get(name string) (*job.Job, bool){
	job, ok := s.jobs[name]
	if !ok {
		return nil, false
	}

	delete(s.jobs, name)
	return job, ok
}

// Peek take a look if the job is in the scheduler queue.
func (s *Runner) Peek(jobName string) bool {
	if _, ok := s.jobs[jobName]; !ok {
		return false
	}
	return true
}

// Bootstrap will assign multiple jobs to the scheduler.
func (s *Runner) Bootstrap(jobs map[string]job.Fn) {
	for name, f := range jobs {
		job := job.NewJob(name, f)
		s.jobs[name] = job
	}
}

