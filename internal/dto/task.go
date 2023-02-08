package dto

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/go-tg-task-bot/internal/entity"
)

func TaskToInlineKb(operation entity.OperationType, buttonText string, task entity.Task) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(buttonText, task.CallbackData(operation).String())
}
