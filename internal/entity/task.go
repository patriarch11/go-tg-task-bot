package entity

import "github.com/gofrs/uuid"

type Task struct {
	ID          uuid.UUID `db:"id" goqu:"skipinsert,skipupdate"`
	SubjectId   uuid.UUID `db:"subject_id" goqu:"skipupdate"`
	Description string    `db:"description"`
}

type ListTasks []*Task

func (t Task) CallbackData(operationType OperationType) Callback {
	return Callback{
		ID:            t.ID,
		OperationType: operationType,
	}
}
