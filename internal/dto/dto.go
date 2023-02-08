package dto

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/go-tg-task-bot/internal/entity"
	"log"
)

func CallbackFromString(str string) entity.Callback {
	var data entity.Callback
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		log.Printf("unmarshaling error: %s", err)
	}
	return data
}

func SubjectToInlineKbRow(operation entity.OperationType, buttonText string, subject entity.Subject) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(buttonText, subject.CallbackData(operation).String())
}

func TaskToInlineKb(operation entity.OperationType, buttonText string, task entity.Task) tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardButtonData(buttonText, task.CallbackData(operation).String())
}
