package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/telegram-task-manager-bot/internal/config"
	"github.com/patriarch11/telegram-task-manager-bot/internal/domain/entity"
)

type MessageService struct {
	adminUserName string
}

func NewMessageService(conf config.BotConfig) *MessageService {
	return &MessageService{adminUserName: conf.Admin}
}

func (s *MessageService) WrapMessageInMainKeyboard(msg tgbotapi.MessageConfig, username string) tgbotapi.MessageConfig {
	if s.adminUserName == username {
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
