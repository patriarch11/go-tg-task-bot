package entity

type Subject struct {
	Id          ID     `db:"id" goqu:"skipinsert,skipupdate" json:"id,omitempty"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

type SubjectList []*Subject
