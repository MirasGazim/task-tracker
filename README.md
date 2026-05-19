# Task Tracker CLI

A simple command line tool to manage your tasks.

## Installation

git clone https://github.com/юзернейм/task-tracker.git
cd task-tracker
go build -o task-cli .

## Usage

### Add a task
task-cli add "Buy groceries"

### List all tasks
task-cli list

### List by status
task-cli list todo
task-cli list done
task-cli list in-progress

### Update a task
task-cli update 1 "Buy milk"

### Delete a task
task-cli delete 1

### Mark as done
task-cli mark-done 1

### Mark as in progress
task-cli mark-in-progress 1