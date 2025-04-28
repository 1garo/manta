package job


type Job struct{
	Name string
	Exec func () error 
}
func NewJob(name string, exec func() error) *Job {
	return &Job{
		Name: name,
		Exec: exec,
	}
}

func (j *Job) Execute() error {
	return j.Exec()
}
