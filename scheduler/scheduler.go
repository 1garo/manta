package scheduler

import "github.com/1garo/manta/job"

type Scheduler struct {
	jobs map[string]*job.Job
}


func NewScheduler() *Scheduler {
	return &Scheduler{
		jobs: make(map[string]*job.Job, 0),
	}
}

func (s *Scheduler) AddJob(job *job.Job) error {
	s.jobs[job.Name] = job
	return nil
}

func (s *Scheduler) Len() int {
	return len(s.jobs)
}

// Get will pop the job from the queue, if found, and return the job and true.
//
// If not, it will return nil and false.
func (s *Scheduler) Get(name string) (*job.Job, bool){
	// TODO: maybe we want double-linked lists here?
	// so we can get and remove the job from the queue on O(1) operation?
	// the idea is that when you get the value, you are gonna pop if from the queue to execute
	// if you want just to see if the job exists, you can use Peek
	job, ok := s.jobs[name]
	if !ok {
		return nil, false
	}

	delete(s.jobs, name)
	return job, ok
}

func (s *Scheduler) Peek(name string) bool {
	if _, ok := s.jobs[name]; !ok {
		return false
	}
	return true
}

// Bootstrap will assign multiple jobs to the scheduler.
//
// TODO: what if we had a function that would improve users's developer experience.
// I am thinking of a function that will call scheduler.AddJob() for multiple jobs.
//
// I believe that this could be a user need at some point. Lets say that I want to call this at my init function
// so I am not manually calling the scheduler.AddJob() a bunch of times.
func (s *Scheduler) Bootstrap(jobs []job.Job) {}

