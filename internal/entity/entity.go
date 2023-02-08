package entity

import "github.com/gofrs/uuid"

type Subject struct {
	ID          uuid.UUID `db:"id" goqu:"skipinsert,skipupdate"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}

type Task struct {
	ID          uuid.UUID `db:"id" goqu:"skipinsert,skipupdate"`
	SubjectId   uuid.UUID `db:"subject_id" goqu:"skipinsert,skipupdate"`
	Description string    `db:"description"`
}

type ListSubjects []*Subject

type ListTasks []*Task
