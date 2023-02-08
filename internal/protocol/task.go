package protocol

import (
	"github.com/gofrs/uuid"
	"github.com/patriarch11/go-tg-task-bot/internal/entity"
)

type PostgresTaskRepository interface {
	Create(task *entity.Task) (*entity.Task, error)
	GetAllBySubjectID(subjectID uuid.UUID) (entity.ListTasks, error)
	Update(task *entity.Task) (*entity.Task, error)
	Delete(id uuid.UUID) error
}
