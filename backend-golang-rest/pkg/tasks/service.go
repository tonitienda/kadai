package tasks

const (
	TaskStatusPending = "pending"
	TaskStatusDone    = "done"
)

type Task struct {
	ID          string
	OwnerID     string
	Title       string
	Description string
	Status      string
}

type TasksDataSource interface {
	GetTasks(ownerId string) ([]Task, error)
	AddTask(task Task) error
	DeleteTask(id string) error
}

type TasksService struct {
	ds TasksDataSource
}

func NewTasksService(ds TasksDataSource) *TasksService {
	return &TasksService{ds: ds}
}

func (s *TasksService) GetTasks(ownerId string) ([]Task, error) {
	return s.ds.GetTasks(ownerId)
}

func (s *TasksService) AddTask(task Task) error {
	return s.ds.AddTask(task)
}

func (s *TasksService) DeleteTask(taskId string) error {
	return s.ds.DeleteTask(taskId)
}
