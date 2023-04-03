package interfaces

import (
	"context"
	"github.com/patriarch11/telegram-task-manager-bot/internal/domain/entity"
)

type TaskStorage interface {
	Create(ctx context.Context, task *entity.Task) (*entity.Task, error)
	Get(ctx context.Context, id entity.ID) (*entity.Task, error)
	GetBySubjectId(ctx context.Context, subjectId entity.ID) (entity.TaskList, error)
	Update(ctx context.Context, task *entity.Task) (*entity.Task, error)
	Delete(ctx context.Context, id entity.ID) error
}

type TaskService interface {
	Create(ctx context.Context, task *entity.Task) (*entity.Task, error)
	Get(ctx context.Context, id entity.ID) (*entity.Task, error)
	GetBySubjectId(ctx context.Context, subjectId entity.ID) (entity.TaskList, error)
	Update(ctx context.Context, task *entity.Task) (*entity.Task, error)
	Delete(ctx context.Context, id entity.ID) error
}
