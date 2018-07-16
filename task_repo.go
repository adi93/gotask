package main

type TaskRepository interface {
	AddTask(*Task) error
	DeleteAll()
	DeleteTask(int)
	GetAll() []*Task
	Validate(*Task) error
}

type InMemoryTaskRepo struct {
	DataStore []*Task
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

func (repo *InMemoryTaskRepo) DeleteTask(id int) {
	deleteIndex := -1
	for i, task := range repo.DataStore {
		if task.Id == id {
			deleteIndex = i
			break
		}
	}
	repo.DataStore = append(repo.DataStore[:deleteIndex], repo.DataStore[deleteIndex+1:]...)
}
