package main

import "time"

type Task struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Status    TaskState `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

func ParseState(input string) (TaskState, bool) {
	for k, v := range stateName {
		if v == input {
			return k, true
		}
	}
	return -1, false
}
