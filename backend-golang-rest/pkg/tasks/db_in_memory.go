package tasks

type InMemoryTasksDB struct {
	tasksByID    map[string]Task
	tasksByOwner map[string][]Task
}

func NewInMemoryTasksDB() *InMemoryTasksDB {
	return &InMemoryTasksDB{
		tasksByID:    make(map[string]Task),
		tasksByOwner: make(map[string][]Task),
	}
}

func (db *InMemoryTasksDB) GetTasks(ownerID string) ([]Task, error) {
	tasks, ok := db.tasksByOwner[ownerID]
	if !ok {
		return []Task{}, nil
	}

	return tasks, nil
}

func (db *InMemoryTasksDB) AddTask(task Task) error {

	db.tasksByID[task.ID] = task
	db.tasksByOwner[task.OwnerID] = append(db.tasksByOwner[task.OwnerID], task)

	return nil
}

func (db *InMemoryTasksDB) DeleteTask(id string) error {
	task := db.tasksByID[id]
	delete(db.tasksByID, id)

	tasks := db.tasksByOwner[task.OwnerID]
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	db.tasksByOwner[task.OwnerID] = tasks
	return nil
}
