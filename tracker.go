package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

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
	case "mark-in-progress":
		changeState(model.InProgress)
	case "mark-done":
		changeState(model.Closed)
	case "list":
		listTask()
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
		ID:        DB.NextID,
		Name:      name,
		CreatedAt: time.Now(),
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

	if error != nil {
		fmt.Println(error)
		return
	}
	task, ok := DB.Tasks[id]
	if !ok {
		fmt.Printf("id:%d not found ", id)
		return
	}

	task.Name = name
	task.UpdatedAt = time.Now()
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

func changeState(state model.TaskState) {
	id, error := strconv.ParseInt(os.Args[2], 0, 64)
	if error != nil {
		fmt.Println(error, "Second parameter should contain id of task")
		return
	}

	DB, task, error := tools.GetTaskByID(id)
	if error != nil {
		fmt.Printf("task with id:%d not found", id)
		return
	}

	task.State = state
	task.UpdatedAt = time.Now()
	DB.Tasks[task.ID] = task
	error = tools.SaveChanges(DB)

	if error != nil {
		println("cannot save changes")
	} else {
		println("Task updated succesfully")
	}
}

func listTask() {
	filter := ""
	if len(os.Args) > 2 {
		filter = os.Args[2]
	}
	DB, error := tools.GetDB()

	if error != nil || DB.Tasks == nil || len(DB.Tasks) == 0 {
		return
	}

	// DB.Tasks.for

	filterFunc := func(model.Task) bool {
		return true
	}

	stateToFilter := -1

	switch filter {
	case "todo":
		stateToFilter = int(model.New)
	case "done":
		stateToFilter = int(model.Closed)
	case "in-progress":
		stateToFilter = int(model.InProgress)
	}

	if stateToFilter != -1 {
		filterFunc = func(task model.Task) bool {
			return task.State == model.TaskState(stateToFilter)
		}
	}

	for _, task := range DB.Tasks {
		if filterFunc(task) {
			printTask(task)
		}
	}
}

func printTask(task model.Task) {
	fmt.Println(task)
}
