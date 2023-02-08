package entity

import "github.com/gofrs/uuid"

type Task struct {
	ID          uuid.UUID `db:"id" goqu:"skipinsert,skipupdate"`
	SubjectId   uuid.UUID `db:"subject_id" goqu:"skipinsert,skipupdate"`
	Description string    `db:"description"`
}

type ListTasks []*Task
