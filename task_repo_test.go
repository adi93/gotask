package main

import (
	"fmt"
	"testing"
)

func TestGetTask(t *testing.T) {
	var taskRepo InMemoryTaskRepo
	taskRepo.AddTask(&Task{Id: 2, Name: "First", Description: "Task"})
	taskRepo.AddTask(&Task{Id: 4, Name: "First", Description: "Task"})
	taskRepo.AddTask(&Task{Id: 1, Name: "First", Description: "Task"})
	task, err := taskRepo.Get(4)
	if err != nil {
		t.Errorf("Did not get task %s", err)
	}
	if task.Id != 4 {
		t.Errorf("Did not get the correct task")
	}
	_, err = taskRepo.Get(3)
	if err == nil {
		t.Errorf("Expected error, found none")
	}

}

func TestGetAll(t *testing.T) {
	var taskRepo InMemoryTaskRepo
	if len(taskRepo.GetAll()) != 0 {
		t.Errorf("Size mismatch, expected 0 tasks, found %d", len(taskRepo.GetAll()))
	}
	taskRepo.AddTask(&Task{Id: 2, Name: "First", Description: "Task"})
	tasks := taskRepo.GetAll()
	if len(tasks) != 1 {
		t.Errorf("Size mismatch, expected 1 tasks, found %d", len(tasks))
	}
	task := tasks[0]
	if task.Id != 2 {
		t.Errorf("Expected task of id 2, found one with id %d", task.Id)
	}
	taskRepo.AddTask(&Task{Id: 3, Name: "First", Description: "Task"})
	taskRepo.AddTask(&Task{Id: 4, Name: "First", Description: "Task"})
	taskRepo.AddTask(&Task{Id: 1, Name: "First", Description: "Task"})
	tasks = taskRepo.GetAll()
	if len(tasks) != 4 {
		t.Errorf("Size mismatch, expected 4 tasks, found %d", len(tasks))
	}
}
func TestAddTask(t *testing.T) {
	var taskRepo InMemoryTaskRepo
	taskRepo.AddTask(&Task{Name: "First", Description: "Task"})
	taskRepo.AddTask(&Task{Name: "Second", Description: "Task"})
	tasks := taskRepo.GetAll()
	if len(tasks) != 2 {
		t.Errorf("Size mismatch, expected 2 tasks, found %d", len(tasks))
	}
	task := tasks[0]
	if task.Name != "First" || task.Description != "Task" {
		t.Errorf("Didn't add task properly.")
	}
	task = tasks[1]
	if task.Name != "Second" || task.Description != "Task" {
		t.Errorf("Didn't add task properly.")
	}

}

func TestDeleteAll(t *testing.T) {
	var taskRepo InMemoryTaskRepo
	taskRepo.AddTask(&Task{Name: "First", Description: "Task"})
	taskRepo.AddTask(&Task{Name: "First", Description: "Task"})
	taskRepo.DeleteAll()
	if len(taskRepo.GetAll()) != 0 {
		t.Errorf("Couldn't delete all tasks")
	}
}

func TestValidate(t *testing.T) {
}

func TestDeleteTask(t *testing.T) {
	var taskRepo InMemoryTaskRepo
	taskRepo.AddTask(&Task{Id: 1, Name: "First", Description: "Task"})
	taskRepo.AddTask(&Task{Id: 2, Name: "Second", Description: "Task"})
	if err := taskRepo.DeleteTask(2); err != nil {
		t.Errorf("Could not delete task")
	}
	if len(taskRepo.GetAll()) != 1 {
		t.Errorf("Error in deletetion, size mismatch")
	}
	_, err := taskRepo.Get(1)
	if err != nil {
		t.Errorf("Deleted the wrong task")
	}
	if err := taskRepo.DeleteTask(3); err == nil {
		t.Errorf("Expected error while deleting non existent task, got none")
	}
	if err := taskRepo.DeleteTask(3); fmt.Sprintf("%s", err) != "No task with given index exists" {
		t.Errorf("Got wrong error message %s", err)
	}

}
