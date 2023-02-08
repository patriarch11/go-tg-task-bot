package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/go-tg-task-bot/internal/dto"
	"github.com/patriarch11/go-tg-task-bot/internal/entity"
)

func (b *Bot) handleCallback(callbackQuery *tgbotapi.CallbackQuery) error {
	c := tgbotapi.NewCallback(callbackQuery.ID, callbackQuery.Data)
	if _, err := b.bot.Request(c); err != nil {
		return err
	}
	callback := dto.CallbackFromString(callbackQuery.Data)

	switch callback.OperationType {
	case entity.GetTasks:
		// get tasks
	case entity.AddTask:
		if !b.isAdmin(callbackQuery.From.UserName) {
			return b.notAdminResponse(callbackQuery.Message.Chat.ID)
		}
		// add task
	case entity.UpdateTask:
		if !b.isAdmin(callbackQuery.From.UserName) {
			return b.notAdminResponse(callbackQuery.Message.Chat.ID)
		}
		// update task
	case entity.DeleteTask:
		if !b.isAdmin(callbackQuery.From.UserName) {
			return b.notAdminResponse(callbackQuery.Message.Chat.ID)
		}
		// delete task
	case entity.UpdateSubject:
		if !b.isAdmin(callbackQuery.From.UserName) {
			return b.notAdminResponse(callbackQuery.Message.Chat.ID)
		}
		// update subject
	case entity.DeleteSubject:
		if !b.isAdmin(callbackQuery.From.UserName) {
			return b.notAdminResponse(callbackQuery.Message.Chat.ID)
		}
		// delete subject
	}
	return nil
}
