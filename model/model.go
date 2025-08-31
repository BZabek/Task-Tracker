// Package model - json object representation
package model

type DB struct {
	Tasks  []Task
	NextID int
}

type Task struct {
	ID    int
	Name  string
	State TaskState
}

type TaskState int

const (
	New TaskState = iota
	InProgress
	Closed
)
