package tasks

import (
	"sync"

	"github.com/tonitienda/kadai/backend-rest-go/pkg/common"
)

type InMemoryTasksDB struct {
	tasksByID    map[string]Task
	tasksByOwner map[string][]Task
	lock         sync.RWMutex
}

func NewInMemoryTasksDB() *InMemoryTasksDB {
	return &InMemoryTasksDB{
		tasksByID:    make(map[string]Task),
		tasksByOwner: make(map[string][]Task),
	}
}

func (db *InMemoryTasksDB) GetTasks(ownerID string) ([]Task, error) {
	db.lock.RLock()
	tasks, ok := db.tasksByOwner[ownerID]
	db.lock.RUnlock()

	if !ok {
		return []Task{}, nil
	}

	notDeletedTasks := []Task{}

	for _, t := range tasks {
		if t.DeletedAt.IsZero() {
			notDeletedTasks = append(notDeletedTasks, t)
		}
	}

	return notDeletedTasks, nil
}

func (db *InMemoryTasksDB) GetTask(taskID string) (Task, bool) {
	db.lock.RLock()
	task, ok := db.tasksByID[taskID]
	db.lock.RUnlock()

	return task, ok
}

func (db *InMemoryTasksDB) AddTask(task Task) error {
	db.lock.Lock()
	db.tasksByID[task.ID] = task
	db.tasksByOwner[task.OwnerID] = append(db.tasksByOwner[task.OwnerID], task)
	db.lock.Unlock()
	return nil
}

func (db *InMemoryTasksDB) UpdateTask(task Task) error {
	db.lock.Lock()
	db.tasksByID[task.ID] = task

	for idx, t := range db.tasksByOwner[task.OwnerID] {
		if t.ID == task.ID {
			db.tasksByOwner[task.OwnerID][idx] = task
		}
	}
	db.lock.Unlock()
	return nil
}

func (db *InMemoryTasksDB) DeleteTask(id string) error {
	db.lock.Lock()
	task := db.tasksByID[id]
	delete(db.tasksByID, id)

	tasks := db.tasksByOwner[task.OwnerID]

	taskIndex := -1
	for i, t := range tasks {
		if t.ID == id {
			taskIndex = i
			break
		}
	}

	if taskIndex == -1 {
		db.lock.Unlock()
		return common.NotFoundError{}
	}

	db.tasksByOwner[task.OwnerID] = append(tasks[:taskIndex], tasks[taskIndex+1:]...)
	db.lock.Unlock()
	return nil
}
