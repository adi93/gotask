package main

import (
	"errors"
)

type TaskRepository interface {
	AddTask(*Task) error
	DeleteAll()
	DeleteTask(int) error
	GetAll() []*Task
	Validate(*Task) error
	Get(int) (*Task, error)
}

type InMemoryTaskRepo struct {
	DataStore []*Task
}

func (repo *InMemoryTaskRepo) Get(taskId int) (task *Task, err error) {
	taskIndex := -1
	for i, task := range repo.DataStore {
		if task.Id == taskId {
			taskIndex = i
			break
		}
	}
	if taskIndex == -1 {
		return nil, errors.New("No task found")
	}
	return repo.DataStore[taskIndex], nil
}

func (repo *InMemoryTaskRepo) AddTask(t *Task) error {
	if err := repo.Validate(t); err != nil {
		return err
	}
	repo.DataStore = append(repo.DataStore, t)
	return nil
}

func (repo *InMemoryTaskRepo) DeleteAll() {
	repo.DataStore = nil
}

func (repo *InMemoryTaskRepo) Validate(task *Task) error {
	return nil
}

func (repo *InMemoryTaskRepo) GetAll() []*Task {
	return repo.DataStore
}

func (repo *InMemoryTaskRepo) DeleteTask(id int) (err error) {
	deleteIndex := -1
	for i, task := range repo.DataStore {
		if task.Id == id {
			deleteIndex = i
			break
		}
	}
	if deleteIndex == -1 {
		return errors.New("No task with given index exists")
	}
	repo.DataStore = append(repo.DataStore[:deleteIndex], repo.DataStore[deleteIndex+1:]...)
	return nil
}
