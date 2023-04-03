package entity

import "fmt"

type Subject struct {
	Id          ID     `db:"id" goqu:"skipinsert,skipupdate" json:"id,omitempty"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

func (s *Subject) MessageFormat() string {
	return fmt.Sprintf("%s\n%s", s.Name, s.Description)
}

func (s *Subject) CallbackData(operationType OperationType) Callback {
	return Callback{
		Id:            s.Id,
		OperationType: operationType,
	}
}

type SubjectList []*Subject
