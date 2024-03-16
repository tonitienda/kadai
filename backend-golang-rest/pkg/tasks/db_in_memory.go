package tasks

import (
	"golang.org/x/exp/maps"
)

type InMemoryTasksDB struct {
	tasks map[string]Task
}

func NewInMemoryTasksDB() *InMemoryTasksDB {
	return &InMemoryTasksDB{}
}

func (db *InMemoryTasksDB) GetTasks() ([]Task, error) {
	return maps.Values(db.tasks), nil
}
