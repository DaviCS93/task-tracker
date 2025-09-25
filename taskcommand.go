package main

type TaskCommand int

const (
	Add TaskCommand = iota
	Update
	Delete
	List
	ListDone
	ListTodo
	ListInProgress
	MarkInProgress
	MarkDone
)

var commandText = map[TaskCommand]string{
	Add:            "add",
	Update:         "update",
	Delete:         "delete",
	List:           "list",
	ListDone:       "list done",
	ListTodo:       "list todo",
	ListInProgress: "list in-progress",
	MarkInProgress: "mark-in-progress",
	MarkDone:       "mark-done",
}

func (ss TaskCommand) String() string {
	return commandText[ss]
}

func ParseCommand(input string) (TaskCommand, bool) {
	for k, v := range commandText {
		if v == input {
			return k, true
		}
	}
	return -1, false
}
