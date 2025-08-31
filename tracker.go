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
	case "delete":
		deleteTask()

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
	if DB.Tasks == nil {
		DB.Tasks = make(map[int64]model.Task)
	}

	DB.Tasks[newTask.ID] = newTask
	DB.NextID = DB.NextID + 1
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

	fmt.Println("123")
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println("123")
	task, ok := DB.Tasks[id]
	if !ok {
		fmt.Printf("id:%d not found ", id)
		return
	}

	task.Name = name
	DB.Tasks[id] = task
	tools.SaveChanges(DB)
}

func deleteTask() {
	id, error := strconv.ParseInt(os.Args[2], 0, 64)
	if error != nil {
		fmt.Println(error, "Second parameter should contain id of task")
		return
	}

	DB, error := tools.GetDB()

	if error != nil {
		fmt.Println(error)
		return
	}

	task, ok := DB.Tasks[id]
	if !ok {
		fmt.Printf("id:%d not found "+task.Name, id)
		return
	}

	delete(DB.Tasks, id)
	tools.SaveChanges(DB)
}
