package entity

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	ShowSubjectsText = "Список предметів"
	AddSubjectText   = "Додати предмет"
)

var (
	ShowSubjectsButton = tgbotapi.NewKeyboardButton(ShowSubjectsText)
	AddSubjectButton   = tgbotapi.NewKeyboardButton(AddSubjectText)
)
