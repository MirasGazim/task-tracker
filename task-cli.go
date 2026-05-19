package main

import (
	"fmt"
	"log"
	"os"
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
			fmt.Println("delete")
		case "update":
			fmt.Println("update")
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
