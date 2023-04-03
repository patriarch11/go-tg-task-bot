package entity

import (
	"database/sql/driver"
	"github.com/gofrs/uuid"
)

type ID struct {
	uuid.UUID
}

func (id ID) Bytes() []byte {
	return id.Bytes()
}

func (id ID) String() string {
	return id.String()
}

func (id ID) Value() (driver.Value, error) {
	return id.UUID.String(), nil
}

func (id *ID) Scan(src interface{}) error {
	return id.UUID.Scan(src)
}

func IDFromString(v string) ID {
	return ID{uuid.FromStringOrNil(v)}
}

type State int

const (
	Default                   State = iota
	ReceiveSubjectName        State = iota
	ReceiveSubjectDescription State = iota
	ReceiveTask               State = iota

	ReceiveUpdSubjectName        State = iota
	ReceiveUpdSubjectDescription State = iota
	ReceiveUpdTask               State = iota
)
