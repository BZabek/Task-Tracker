package main

import (
	"fmt"
	"os"
	"strings"

	"bzabek/task-tracker/model"
	"bzabek/task-tracker/tools"
)

func main() {
	tools.CreateFileIfNotExist()
	firstArg := os.Args[1]
	fmt.Println(firstArg)

	switch firstArg {
	case "add":
		add()
	}
}

func add() {
	name := strings.Join(os.Args[2:], "")

	DB, error := tools.GetDB()
	if error != nil {
		fmt.Println("Some error occured while adding a task")
		return
	}

	newTask := model.Task{
		ID:   DB.NextID,
		Name: name,
	}

	DB.NextID = DB.NextID + 1
	DB.Tasks = append(DB.Tasks, newTask)

	error = tools.SaveChanges(DB)

	if error != nil {
		fmt.Println("Some error occured while adding a task")
	} else {
		fmt.Printf("Task added successfully (ID: %d)", newTask.ID)
	}
}
