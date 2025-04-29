package job

type Job struct {
	Name string
	exec func() error
}

func NewJob(name string, f func() error) *Job {
	return &Job{
		Name: name,
		exec: f,
	}
}

func (j *Job) Execute() error {
	return j.exec()
}
