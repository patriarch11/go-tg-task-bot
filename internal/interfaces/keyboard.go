package interfaces

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/telegram-task-manager-bot/internal/domain/entity"
)

type KeyboardService interface {
	WrapMessageInMainKeyboard(msg tgbotapi.MessageConfig, isAdmin bool) tgbotapi.MessageConfig
	WrapMessageInCancelKeyboard(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig
	WrapSubjectMessageInInlineKeyboard(msg tgbotapi.MessageConfig, subject *entity.Subject, isAdmin bool) tgbotapi.MessageConfig
	WrapTaskMessageInInlineKeyboard(msg tgbotapi.MessageConfig, task *entity.Task, isAdmin bool) tgbotapi.MessageConfig
}
