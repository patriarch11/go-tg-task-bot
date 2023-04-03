package service

import (
	"context"
	"github.com/patriarch11/go-tg-task-bot/internal/domain/entity"
	"github.com/patriarch11/go-tg-task-bot/internal/interfaces"
)

type TaskService struct {
	storage interfaces.TaskStorage
}

func NewTaskService(repository interfaces.TaskStorage) *TaskService {
	return &TaskService{repository}
}

func (s *TaskService) Create(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	return s.storage.Create(ctx, task)
}
func (s *TaskService) Get(ctx context.Context, id entity.ID) (*entity.Task, error) {
	return s.storage.Get(ctx, id)
}

func (s *TaskService) GetBySubjectId(ctx context.Context, subjectId entity.ID) (entity.TaskList, error) {
	return s.storage.GetBySubjectId(ctx, subjectId)
}

func (s *TaskService) Update(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	return s.storage.Update(ctx, task)
}

func (s *TaskService) Delete(ctx context.Context, id entity.ID) error {
	return s.storage.Delete(ctx, id)
}
