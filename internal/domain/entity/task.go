package entity

type Task struct {
	Id          ID     `db:"id" goqu:"skipinsert,skipupdate" json:"id,omitempty"`
	SubjectId   ID     `db:"subject_id" goqu:"skipupdate" json:"subject_id"`
	Description string `db:"description" json:"description"`
}

type TaskList []*Task
