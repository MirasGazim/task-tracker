package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lenofargs := len(os.Args)
	if lenofargs < 2 {
		log.Fatal("Usage: task-cli <command>")
	} else {
		command := os.Args[1]
		switch command {
		case "add":
			err := AddTask(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Task added successfully")
		case "delete":
			change, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal("error: string int")
			}
			err = deleteTask(change)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Task deleted successfully")
		case "update":
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal("error: invalid id")
			}
			err = updateTask(id, os.Args[3])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Task updated successfully")
		case "list":
			listTasks()
		case "mark-done":
			fmt.Println("mark-done")
		case "mark-in-progress":
			fmt.Println("mark-in-progress")
		default:
			fmt.Println("unknown command")
		}
	}
}
