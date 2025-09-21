package main

import "time"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      TaskState `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskState int

const (
	ToDo TaskState = iota
	InProgress
	Done
)

var stateName = map[TaskState]string{
	ToDo:       "to-do",
	InProgress: "in-progress",
	Done:       "done",
}

func (ss TaskState) String() string {
	return stateName[ss]
}
