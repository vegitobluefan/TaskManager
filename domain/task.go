package domain

type Task struct {
	ID      string
	Type    string
	Payload string
	Status  string
	Result  string
}

type TaskRepository interface {
	Save(task *Task) error
	UpdateStatus(id string, status string, result string) error
	GetByID(id string) (*Task, error)
	List() ([]*Task, error)
}
