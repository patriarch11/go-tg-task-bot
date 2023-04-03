package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/telegram-task-manager-bot/internal/domain/entity"
)

type KeyboardService struct {
}

func NewKeyboardService() *KeyboardService {
	return &KeyboardService{}
}

func (s *KeyboardService) WrapMessageInMainKeyboard(msg tgbotapi.MessageConfig, isAdmin bool) tgbotapi.MessageConfig {
	if isAdmin {
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(entity.ShowSubjectsButton, entity.AddSubjectButton),
		)
		return msg
	}
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(entity.ShowSubjectsButton),
	)
	return msg
}

func (s *KeyboardService) WrapMessageInCancelKeyboard(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig {
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(entity.CancelButton),
	)
	return msg
}

func (s *KeyboardService) WrapSubjectMessageInInlineKeyboard(
	msg tgbotapi.MessageConfig, subject *entity.Subject, isAdmin bool) tgbotapi.MessageConfig {
	if isAdmin {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				s.buttonFromSubject(entity.ShowTasks, entity.ShowInlineButtonText, subject),
				s.buttonFromSubject(entity.AddTask, entity.AddInlineButtonText, subject),
			),
			tgbotapi.NewInlineKeyboardRow(
				s.buttonFromSubject(entity.UpdateSubject, entity.UpdateInlineButtonText, subject),
				s.buttonFromSubject(entity.DeleteSubject, entity.DeleteInlineButtonText, subject),
			),
		)
		return msg
	}
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			s.buttonFromSubject(entity.ShowTasks, entity.ShowInlineButtonText, subject),
		),
	)
	return msg
}

func (s *KeyboardService) WrapTaskMessageInInlineKeyboard(
	msg tgbotapi.MessageConfig, task *entity.Task, isAdmin bool) tgbotapi.MessageConfig {
	if isAdmin {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				s.buttonFromTask(entity.UpdateTask, entity.UpdateInlineButtonText, task),
				s.buttonFromTask(entity.DeleteTask, entity.DeleteInlineButtonText, task),
			),
		)
	}
	return msg
}

func (s *KeyboardService) buttonFromSubject(operation entity.OperationType,
	text string, subject *entity.Subject) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(text, subject.CallbackData(operation).String())
}

func (s *KeyboardService) buttonFromTask(operation entity.OperationType,
	text string, task *entity.Task) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(text, task.CallbackData(operation).String())
}
