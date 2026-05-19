package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

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

func saveTasks(tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return errors.New("can not marshall tasks")
	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return errors.New("can not write on file")
	}
	return nil
}

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

func listTasks() {
	tasks := loadTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	for i := range tasks {
		fmt.Printf("ID: %d | %s | %s\n", tasks[i].ID, tasks[i].Description, tasks[i].Status)
	}
}

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
