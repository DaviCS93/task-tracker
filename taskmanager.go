package main

import (
	"encoding/json"
	"os"
	"time"
)

const fileName = "tasks.json"

func CreateTask(name string) (Task, error) {

	var tasks []Task
	tasks, err := readtasks()

	if err != nil {
		return Task{}, err
	}

	// Create new task
	newID := len(tasks) + 1 // simple incremental ID
	newTask := Task{
		ID:        newID,
		Name:      name,
		Status:    ToDo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	// Append and save
	tasks = append(tasks, newTask)
	if err := savetask(tasks); err != nil {
		return Task{}, err
	}

	return newTask, nil

}

func DeleteTask(id int) error {
	var tasks []Task
	tasks, err := readtasks()
	if err != nil {
		return err
	}

	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	if err := savetask(tasks); err != nil {
		return err
	}
	return nil
}

func UpdateTaskName(id int, name string) (Task, error) {
	var tasks []Task
	var updatedTask Task
	tasks, err := readtasks()
	if err != nil {
		return Task{}, err
	}
	for i := 0; i < len(tasks); i++ {

		if tasks[i].ID == id {
			tasks[i].Name = name
			updatedTask = tasks[i]
		}
	}
	if err := savetask(tasks); err != nil {
		return Task{}, err
	}
	return updatedTask, nil
}

func MarkTaskInProgress(id int) error {
	var tasks []Task
	tasks, err := readtasks()
	if err != nil {
		return err
	}
	for i := 0; i < len(tasks); i++ {

		if tasks[i].ID == id {
			tasks[i].Status = InProgress
		}
	}
	if err := savetask(tasks); err != nil {
		return err
	}
	return nil
}

func MarkTaskDone(id int) error {
	tasks := check(readtasks())
	for i := 0; i < len(tasks); i++ {

		if tasks[i].ID == id {
			tasks[i].Status = TaskState(Done)
		}
	}
	if err := savetask(tasks); err != nil {
		return err
	}
	return nil
}

func ListTasks() ([]Task, error) {
	tasks, err := readtasks()
	if err != nil {
		return []Task{}, err
	}
	return tasks, nil
}

func ListTasksByStatus(state TaskState) ([]Task, error) {
	var tasks []Task
	tasks, err := readtasks()
	if err != nil {
		return []Task{}, err
	}
	var taskNamesFound []Task
	for _, t := range tasks {

		if t.Status == state {
			taskNamesFound = append(taskNamesFound, t)
		}
	}
	return taskNamesFound, nil
}

func readtasks() ([]Task, error) {
	// Read file
	var tasks []Task
	file := check(os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666))
	defer file.Close()

	// Check for tasks in the file
	fileInfo := check(file.Stat())
	if fileInfo.Size() > 0 {
		if err := json.NewDecoder(file).Decode(&tasks); err != nil {
			return nil, err
		}
	}
	return tasks, nil
}

func savetask(tasks []Task) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	if err := file.Truncate(0); err != nil {
		return err
	}
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}
	if err := json.NewEncoder(file).Encode(tasks); err != nil {
		return err
	}

	return nil
}
