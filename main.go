package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cmd := check(startCLIMenu())

	switch cmd {
	case Add:
		name := parts[1]
		task := check(CreateTask(name))
		if (Task{}) != task {
			fmt.Printf("Created task: %s\n", task.Name)
		}

	case Update:
		id := check(parseId(parts[1]))
		name := parts[2]
		task := check(UpdateTaskName(id, name))
		if (Task{}) != task {
			fmt.Printf("Updated task: %s\n", task.Name)
		}

	case Delete:
		id := check(parseId(parts[1]))
		if err := DeleteTask(id); err != nil {
			fmt.Printf("Task %d deleted\n", id)
		}

	case List:
		var status TaskState = 99
		if len(parts) > 1 {
			status, _ = ParseState(parts[1])
		}
		var tasks []Task
		switch status {
		case Done:
			tasks = check(ListTasksByStatus(Done))
		case ToDo:
			tasks = check(ListTasksByStatus(ToDo))
		case InProgress:
			tasks = check(ListTasksByStatus(InProgress))
		default:
			tasks = check(ListTasks())
		}
		printTasks(tasks)

	case MarkInProgress:
		id := check(parseId(parts[1]))
		if err := MarkTaskInProgress(id); err == nil {
			fmt.Printf("Task %d updated to \"In progress\"\n", id)
		} else {
			fmt.Printf("Error marking requested task: %s\n", err.Error())
		}

	case MarkDone:
		id := check(parseId(parts[1]))
		if err := MarkTaskDone(id); err == nil {
			fmt.Printf("Task %d updated to \"Done\"\n", id)
		} else {
			fmt.Printf("Error marking requested task: %s\n", err.Error())

		}

	default:
		fmt.Printf("Unknown command: %s\n", cmd)
	}
}

func printTasks(tasks []Task) {
	fmt.Printf("Tasks found: %d\n", len(tasks))
	for _, t := range tasks {
		fmt.Printf("[%d] %s - %s (Created: %s, Updated: %s)\n",
			t.ID, t.Name, t.Status.String(),
			t.CreatedAt.Format("2006-01-02 15:04"),
			t.UpdatedAt.Format("2006-01-02 15:04"))
	}
}

func parseId(commandpart string) (Task, error) {
	idInt, err := strconv.Atoi(commandpart)
	if err != nil {
		return 0, err
	}
	return idInt, nil
}

func startCLIMenu() (TaskCommand, error) {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Task Tracker CLI")
	fmt.Println(`Commands:
	add <taskName>
	update <taskId> <new taskName>
	delete <taskId>
	list
	list done
	list todo
	list in-progress
	mark-in-progress <taskId>
	mark-done <taskId>`)
	for {
		fmt.Print("> ")
		if !reader.Scan() {
			break
		}

		input := strings.TrimSpace(reader.Text())
		if input == "" {
			continue
		}

		parts := strings.SplitN(input, " ", 3)
		cmdStr := parts[0]
		cmd, ok := ParseCommand(cmdStr)
		if !ok {
			fmt.Printf("Unknown command: %s\n", cmdStr)
			continue
		}
	}
}
