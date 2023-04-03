package interfaces

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/go-tg-task-bot/internal/domain/entity"
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

type UseCaseTask interface {
	ShowTasks(handler UpdateHandler, subjectId entity.ID, chatId int64, isAdmin bool) error
	AddTaskReply(handler UpdateHandler, subjectId entity.ID, chatId int64) error
	ReceiveTaskDescriptionAndSave(handler UpdateHandler, msg *tgbotapi.Message) error
	UpdateTaskReply(handler UpdateHandler, taskId entity.ID, chatId int64) error
	ReceiveUpdTaskDescriptionAndSave(handler UpdateHandler, msg *tgbotapi.Message) error
	DeleteTask(handler UpdateHandler, taskId entity.ID, chatId int64) error
}
