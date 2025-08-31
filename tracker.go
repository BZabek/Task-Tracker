package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"bzabek/task-tracker/model"
	"bzabek/task-tracker/tools"
)

func main() {
	tools.CreateFileIfNotExist()
	firstArg := os.Args[1]

	switch firstArg {
	case "add":
		add()
	case "update":
		update()
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

func update() {
	id, error := strconv.ParseInt(os.Args[2], 0, 64)
	if error != nil {
		fmt.Println(error, "Second parameter should contain id of task")
		return
	}

	name := strings.Join(os.Args[3:], "")

	if len(name) == 0 {
		fmt.Println("incorect name")
		return
	}

	DB, error := tools.GetDB()
	if error != nil {
		fmt.Println("Some error occured while adding a task")
		return
	}

	if DB.Tasks != nil {
		for i := range DB.Tasks {
			if DB.Tasks[i].ID == id {
				DB.Tasks[i].Name = name
				tools.SaveChanges(DB)
				return
			}
		}
	}

	fmt.Printf("id:%d not found ", id)
}
