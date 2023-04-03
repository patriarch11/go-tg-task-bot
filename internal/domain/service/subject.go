package service

import (
	"context"
	"github.com/patriarch11/telegram-task-manager-bot/internal/domain/entity"
	"github.com/patriarch11/telegram-task-manager-bot/internal/interfaces"
)

type SubjectService struct {
	storage interfaces.SubjectStorage
}

func NewSubjectService(repository interfaces.SubjectStorage) *SubjectService {
	return &SubjectService{repository}
}

func (s *SubjectService) Create(ctx context.Context, subject *entity.Subject) (*entity.Subject, error) {
	return s.storage.Create(ctx, subject)
}

func (s *SubjectService) Get(ctx context.Context, id entity.ID) (*entity.Subject, error) {
	return s.storage.Get(ctx, id)
}

func (s *SubjectService) GetAll(ctx context.Context) (entity.SubjectList, error) {
	return s.storage.GetAll(ctx)
}

func (s *SubjectService) Update(ctx context.Context, subject *entity.Subject) (*entity.Subject, error) {
	return s.storage.Update(ctx, subject)
}

func (s *SubjectService) Delete(ctx context.Context, id entity.ID) error {
	return s.storage.Delete(ctx, id)
}
