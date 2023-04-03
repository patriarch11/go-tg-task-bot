package entity

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	ShowSubjectsButtonText = "Список предметів"
	AddSubjectButtonText   = "Додати предмет"
	CancelButtonText       = "Відмінити"
	ShowInlineButtonText   = "📋"
	AddInlineButtonText    = "➕"
	DeleteInlineButtonText = "➖"
	UpdateInlineButtonText = "📝"
)

var (
	ShowSubjectsButton = tgbotapi.NewKeyboardButton(ShowSubjectsButtonText)
	AddSubjectButton   = tgbotapi.NewKeyboardButton(AddSubjectButtonText)
	CancelButton       = tgbotapi.NewKeyboardButton(CancelButtonText)
)
