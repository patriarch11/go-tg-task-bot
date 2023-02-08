package entity

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"log"
)

type Subject struct {
	ID          uuid.UUID `db:"id" goqu:"skipinsert,skipupdate"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}

type ListSubjects []*Subject

type SubjectCallback struct {
	ID            uuid.UUID     `json:"id"`
	OperationType OperationType `json:"operation_type"`
}

func (s Subject) CallbackData(operationType OperationType) SubjectCallback {
	return SubjectCallback{
		ID:            s.ID,
		OperationType: operationType,
	}
}

func (s SubjectCallback) String() string {
	str, _ := json.Marshal(s)
	log.Printf("marshaled data: %s", string(str))
	return string(str)
}
