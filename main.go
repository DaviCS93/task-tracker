package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

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

		parts := strings.SplitN(input, " ", 2)
		cmdStr := parts[0]
		cmd, ok := ParseCommand(cmdStr)
		if !ok {
			fmt.Printf("Unknown command: %s", cmdStr)
			continue
		}

		switch cmd {
		case Add:
			name := parts[1]
			task := check(CreateTask(name))
			if (Task{}) != task {
				fmt.Printf("Created task: %s", task.Name)
			}
		case Update:
			id := check(parseId(parts[1]))
			name := parts[2]
			task := check(UpdateTaskName(id, name))
			if (Task{}) != task {
				fmt.Printf("Created task: %s", task.Name)
			}
		case Delete:
			id := check(parseId(parts[1]))
			if err := DeleteTask(id); err != nil {
				fmt.Printf("Task %d deleted", id)
			}

		case List:
			status, _ := ParseState(parts[1])
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
			if err := MarkTaskInProgress(id); err != nil {
				fmt.Printf("Task %d updated to \"In progress\"", id)
			}
		case MarkDone:
			id := check(parseId(parts[1]))
			if err := MarkTaskDone(id); err != nil {
				fmt.Printf("Task %d updated to \"Done\"", id)
			}
		default:
			fmt.Println("Unknown command:", cmd)
		}
	}

}

func printTasks(tasks []Task) {
	for _, t := range tasks {
		fmt.Printf("[%d] %s - %s (Created: %s, Updated: %s)\n",
			t.ID, t.Name, t.Status.String(),
			t.CreatedAt.Format("2006-01-02 15:04"),
			t.UpdatedAt.Format("2006-01-02 15:04"))
	}
}

func parseId(commandpart string) (int, error) {
	idInt, err := strconv.Atoi(commandpart)
	if err != nil {
		return 0, err
	}
	return idInt, nil
}
