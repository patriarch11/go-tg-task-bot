package entity

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type OperationType int

const (
	ShowTasks     OperationType = iota
	AddTask       OperationType = iota
	UpdateSubject OperationType = iota
	DeleteSubject OperationType = iota

	UpdateTask OperationType = iota
	DeleteTask OperationType = iota
)

type Callback struct {
	Id            ID            `json:"id"`
	OperationType OperationType `json:"operation_type"`
}

func (s Callback) String() string {
	str, _ := json.Marshal(s)
	logrus.Infof("marshaled data: %s", string(str))
	return string(str)
}
