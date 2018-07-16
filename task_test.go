package main

import (
	"testing"
)

func TestCompleteTas(t *testing.T) {
	task := &Task{Name: "First task", Description: "A task made for example", Priority: Medium}
	task.CompleteTask()
	if task.Status != Done {
		t.Errorf("Did not change status")
	}
}

func TestString(t *testing.T) {
	task := &Task{Name: "First task", Description: "A task made for example", Priority: Medium}
	result := "Task: First task, Description: A task made for example"
	if result != task.String() {
		t.Errorf("Expected \"%s\", got \"%s\"", result, task.String())
	}
}
