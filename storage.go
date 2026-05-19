package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

// loadTasks reads tasks from Tasks.json and returns a slice of Task.
// If the file does not exist, it returns an empty slice.

func loadTasks() []Task {

	data, err := os.ReadFile("Tasks.json")
	if err != nil {
		return []Task{}
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal()
	}
	return tasks
}

// saveTasks takes a slice of Task and writes it to tasks.json.
// Returns an error if marshaling or writing fails.

func saveTasks(tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return errors.New("can not marshall tasks")
	}
	err = os.WriteFile("Tasks.json", data, 0644)
	if err != nil {
		return errors.New("can not write on file")
	}
	return nil
}

// this function add task to Tasks.json

func AddTask(description string) error {
	tasks := loadTasks()
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	new_task := Task{
		ID:          maxID + 1,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, new_task)
	err := saveTasks(tasks)
	if err != nil {
		return err
	}
	return nil
}

// listTasks prints all tasks to the terminal.
// You can also filter tasks by status: "todo", "done", "in-progress".
// If filter is empty, all tasks are printed.

func listTasks(filter string) {
	tasks := loadTasks()
	for _, task := range tasks {
		if filter == "" || task.Status == filter {
			fmt.Printf("ID: %d | %s | %s\n", task.ID, task.Description, task.Status)
		}
	}
}

// deleteTask removes the task with the given ID from tasks.json.
// Returns an error if the task is not found.

func deleteTask(id int) error {
	tasks := loadTasks()
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks(tasks)
			return nil
		}
	}
	return errors.New("Not found task")
}

// updateTask updates the description of the task with the given ID.
// Returns an error if the task is not found.

func updateTask(id int, description string) error {
	tasks := loadTasks()
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			saveTasks(tasks)
			return nil
		}
	}
	return errors.New("not found task to update")
}

// markDone sets the status of the task with the given ID to "done".
// Returns an error if the task is not found or already marked as done.

func markDone(id int) error {
	tasks := loadTasks()
	for i, task := range tasks {
		if task.ID == id {
			if tasks[i].Status == "done" {
				return errors.New("status already mark-done")
			} else {
				tasks[i].Status = "done"
				tasks[i].UpdatedAt = time.Now()
				saveTasks(tasks)
				return nil
			}
		}
	}
	return errors.New("didn't find task")
}

// markInProgress sets the status of the task with the given ID to "in-progress".
// Returns an error if the task is not found or already marked as in-progress.

func markInProgress(id int) error {
	tasks := loadTasks()
	for i, task := range tasks {
		if task.ID == id {
			if tasks[i].Status == "in-progress" {
				return errors.New("status already in-progress")
			} else {
				tasks[i].Status = "in-progress"
				tasks[i].UpdatedAt = time.Now()
				saveTasks(tasks)
				return nil
			}
		}
	}
	return errors.New("didn't find task")
}
