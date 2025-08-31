// Package model - json object representation
package model

type DB struct {
	Tasks  []Task
	NextID int64
}

type Task struct {
	ID    int64
	Name  string
	State TaskState
}

type TaskState int

const (
	New TaskState = iota
	InProgress
	Closed
)
