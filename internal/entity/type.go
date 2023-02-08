package entity

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"log"
)

type OperationType int

const (
	GetTasks OperationType = iota
	UpdateSubject
	DeleteSubject
	AddTask
	UpdateTask
	DeleteTask
)

type Callback struct {
	ID            uuid.UUID     `json:"id"`
	OperationType OperationType `json:"operation_type"`
}

func (s Callback) String() string {
	str, _ := json.Marshal(s)
	log.Printf("marshaled data: %s", string(str))
	return string(str)
}
