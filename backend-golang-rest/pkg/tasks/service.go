package tasks

const (
	TaskStatusPending = "pending"
	TaskStatusDone    = "done"
)

type Task struct {
	ID          string
	Title       string
	Description string
	Status      string
}

type TasksDataSource interface {
	GetTasks() ([]Task, error)
}

type TasksService struct {
	ds TasksDataSource
}

func NewTasksService(ds TasksDataSource) *TasksService {
	return &TasksService{ds: ds}
}

func (s *TasksService) GetTasks() ([]Task, error) {
	return s.ds.GetTasks()
}
