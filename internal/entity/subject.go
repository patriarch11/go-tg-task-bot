package entity

import (
	"github.com/gofrs/uuid"
)

type Subject struct {
	ID          uuid.UUID `db:"id" goqu:"skipinsert,skipupdate"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}

type ListSubjects []*Subject

func (s Subject) CallbackData(operationType OperationType) Callback {
	return Callback{
		ID:            s.ID,
		OperationType: operationType,
	}
}
