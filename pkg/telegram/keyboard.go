package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var subjects = map[string]string{
	"geodezia": "id1",
	"zalupa":   "id2",
}

func NewKb() []tgbotapi.InlineKeyboardButton {
	kb := tgbotapi.NewInlineKeyboardRow()
	for key, value := range subjects {
		kb = append(kb,
			tgbotapi.NewInlineKeyboardButtonData(key, value),
		)
	}
	return kb
}

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)
