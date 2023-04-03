package subject

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/go-tg-task-bot/internal/domain/entity"
	"github.com/patriarch11/go-tg-task-bot/internal/interfaces"
)

type UseCaseSubject struct {
	subject         *entity.Subject
	subjectService  interfaces.SubjectService
	keyboardService interfaces.KeyboardService
}

func NewUseCaseSubject(
	subjectService interfaces.SubjectService,
	keyboardService interfaces.KeyboardService) *UseCaseSubject {
	return &UseCaseSubject{
		subject:         &entity.Subject{},
		subjectService:  subjectService,
		keyboardService: keyboardService,
	}
}

func (u *UseCaseSubject) ShowAllSubjects(handler interfaces.UpdateHandler, msg *tgbotapi.Message, isAdmin bool) error {
	subjects, err := u.subjectService.GetAll(context.Background())
	if err != nil {
		return err
	}

	for _, subject := range subjects {
		rep := tgbotapi.NewMessage(msg.Chat.ID, subject.MessageFormat())
		rep = u.keyboardService.WrapSubjectMessageInInlineKeyboard(rep, subject, isAdmin)
		_, err = handler.BotAPI().Send(rep)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *UseCaseSubject) AddSubjectReply(handler interfaces.UpdateHandler, msg *tgbotapi.Message) error {
	rep := tgbotapi.NewMessage(msg.Chat.ID, entity.InputSubjectNameReply)
	rep = u.keyboardService.WrapMessageInCancelKeyboard(rep)

	handler.SetState(entity.ReceiveSubjectName)

	_, err := handler.BotAPI().Send(rep)
	if err != nil {
		handler.SetState(entity.Default)
		return err
	}
	return nil
}

func (u *UseCaseSubject) SetSubjectName(handler interfaces.UpdateHandler, msg *tgbotapi.Message) error {
	u.subject.Name = msg.Text

	rep := tgbotapi.NewMessage(msg.Chat.ID, entity.InputSubjectDescriptionReply)
	rep = u.keyboardService.WrapMessageInCancelKeyboard(rep)

	handler.SetState(entity.ReceiveSubjectDescription)

	_, err := handler.BotAPI().Send(rep)
	if err != nil {
		handler.SetState(entity.Default)
		return err
	}
	return nil
}

func (u *UseCaseSubject) ReceiveSubjectDescriptionAndSave(handler interfaces.UpdateHandler, msg *tgbotapi.Message) error {
	u.subject.Description = msg.Text

	handler.SetState(entity.Default)

	_, err := u.subjectService.Create(context.Background(), u.subject)
	if err != nil {
		return err
	}

	rep := tgbotapi.NewMessage(msg.Chat.ID, entity.SubjectAddedReply)
	rep = u.keyboardService.WrapMessageInMainKeyboard(rep, true)

	_, err = handler.BotAPI().Send(rep)
	if err != nil {
		u.subject = &entity.Subject{}
		handler.SetState(entity.Default)
		return err
	}
	u.subject = &entity.Subject{}
	return nil
}

func (u *UseCaseSubject) UpdateSubjectReply(handler interfaces.UpdateHandler, subjectId entity.ID, chatId int64) error {
	rep := tgbotapi.NewMessage(chatId, entity.InputSubjectNameReply)
	rep = u.keyboardService.WrapMessageInCancelKeyboard(rep)
	u.subject.Id = subjectId

	handler.SetState(entity.ReceiveUpdSubjectName)

	_, err := handler.BotAPI().Send(rep)
	if err != nil {
		u.subject = &entity.Subject{}
		handler.SetState(entity.Default)
		return err
	}
	return nil
}

func (u *UseCaseSubject) SetUpdSubjectName(handler interfaces.UpdateHandler, msg *tgbotapi.Message) error {
	u.subject.Name = msg.Text

	rep := tgbotapi.NewMessage(msg.Chat.ID, entity.InputSubjectDescriptionReply)
	rep = u.keyboardService.WrapMessageInCancelKeyboard(rep)

	handler.SetState(entity.ReceiveUpdSubjectDescription)

	_, err := handler.BotAPI().Send(rep)
	if err != nil {
		u.subject = &entity.Subject{}
		handler.SetState(entity.Default)
		return err
	}
	return nil
}

func (u *UseCaseSubject) ReceiveUpdSubjectDescriptionAndSave(handler interfaces.UpdateHandler, msg *tgbotapi.Message) error {
	u.subject.Description = msg.Text

	handler.SetState(entity.Default)

	_, err := u.subjectService.Update(context.Background(), u.subject)
	if err != nil {
		u.subject = &entity.Subject{}
		handler.SetState(entity.Default)
		return err
	}

	rep := tgbotapi.NewMessage(msg.Chat.ID, entity.SubjectUpdatedReply)
	rep = u.keyboardService.WrapMessageInMainKeyboard(rep, true)

	_, err = handler.BotAPI().Send(rep)
	if err != nil {
		u.subject = &entity.Subject{}
		handler.SetState(entity.Default)
		return err
	}
	u.subject = &entity.Subject{}
	return nil
}

func (u *UseCaseSubject) DeleteSubject(handler interfaces.UpdateHandler, subjectId entity.ID, chatId int64) error {
	err := u.subjectService.Delete(context.Background(), subjectId)
	if err != nil {
		return err
	}
	rep := tgbotapi.NewMessage(chatId, entity.SubjectDeletedReply)
	rep = u.keyboardService.WrapMessageInMainKeyboard(rep, true)
	_, err = handler.BotAPI().Send(rep)
	if err != nil {
		return err
	}
	return nil
}
