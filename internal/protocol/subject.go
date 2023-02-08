package protocol

import (
	"github.com/gofrs/uuid"
	"github.com/patriarch11/go-tg-task-bot/internal/entity"
)

type PostgresSubjectRepository interface {
	Create(subject *entity.Subject) (*entity.Subject, error)
	GetAll() (entity.ListSubjects, error)
	Update(id uuid.UUID, subject *entity.Subject) (*entity.Subject, error)
	Delete(id uuid.UUID) error
}
