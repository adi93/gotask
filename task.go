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

type Story struct {
	Id                int    `json:"id"`
	Name, Description string `json:"name"`
	Tasks             []Task `json:"tasks"`
}

type Task struct {
	Id                int       `json:"id"`
	Name, Description string    `json:"name"`
	Status            Status    `json:"status"`
	Created           time.Time `json:"created"`
	Priority          Priority  `json:"priority"`
	LinkedTasks       []Task    `json:"linked_tasks"`
}

func (task *Task) String() string {
	return fmt.Sprintf("Task: %s, Description: %s", task.Name, task.Description)
}

func (task *Task) CompleteTask() {
	task.Status = Done
}
