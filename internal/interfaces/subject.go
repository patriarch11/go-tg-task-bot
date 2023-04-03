package interfaces

import (
	"context"
	"github.com/patriarch11/telegram-task-manager-bot/internal/domain/entity"
)

type SubjectService interface {
	Create(ctx context.Context, subject *entity.Subject) (*entity.Subject, error)
	Get(ctx context.Context, id entity.ID) (*entity.Subject, error)
	GetAll(ctx context.Context) (entity.SubjectList, error)
	Update(ctx context.Context, subject *entity.Subject) (*entity.Subject, error)
	Delete(ctx context.Context, id entity.ID) error
}

type SubjectStorage interface {
	Create(ctx context.Context, subject *entity.Subject) (*entity.Subject, error)
	Get(ctx context.Context, id entity.ID) (*entity.Subject, error)
	GetAll(ctx context.Context) (entity.SubjectList, error)
	Update(ctx context.Context, subject *entity.Subject) (*entity.Subject, error)
	Delete(ctx context.Context, id entity.ID) error
}