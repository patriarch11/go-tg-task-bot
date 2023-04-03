package interfaces

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

type UseCaseSubject interface {
	ShowAllSubjects(handler UpdateHandler, msg *tgbotapi.Message, isAdmin bool) error
	AddSubjectReply(handler UpdateHandler, msg *tgbotapi.Message) error
	SetSubjectName(handler UpdateHandler, msg *tgbotapi.Message) error
	ReceiveSubjectDescriptionAndSave(handler UpdateHandler, msg *tgbotapi.Message) error
	UpdateSubjectReply(handler UpdateHandler, subjectId entity.ID, chatId int64) error
	SetUpdSubjectName(handler UpdateHandler, msg *tgbotapi.Message) error
	ReceiveUpdSubjectDescriptionAndSave(handler UpdateHandler, msg *tgbotapi.Message) error
	DeleteSubject(handler UpdateHandler, subjectId entity.ID, chatId int64) error
}
