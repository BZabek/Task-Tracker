// Package model - json object representation
package model

type DB struct {
	NextID int64
	Tasks  map[int64]Task
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
