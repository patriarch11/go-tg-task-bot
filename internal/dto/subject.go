package dto

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/go-tg-task-bot/internal/entity"
)

func SubjectToInlineKbRow(operation entity.OperationType, buttonText string, subject entity.Subject) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(buttonText, subject.CallbackData(operation).String())
}
