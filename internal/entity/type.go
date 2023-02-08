package entity

type OperationType int

const (
	GetTasks OperationType = iota
	UpdateSubject
	DeleteSubject
	AddTask
	UpdateTask
	DeleteTask
)
