package main

import (
	"fmt"
	"time"
)

type Status int

const (
	Created Status = iota
	Picked
	Progressing
	Done
)

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

type Task struct {
	Id                int
	Name, Description string
	Status            Status
	Created           time.Time
	Priority          Priority
	LinkedTasks       []Task
}

func (task *Task) String() string {
	return fmt.Sprintf("Task: %s, Description: %s", task.Name, task.Description)
}

func (task *Task) CompleteTask() {
	task.Status = Done
}
