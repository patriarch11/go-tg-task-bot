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
		return b.showTasks(callbackQuery.Message.Chat.ID, callback.ID)
	case entity.AddTask:
		if !b.isAdmin(callbackQuery.From.UserName) {
			return b.notAdminResponse(callbackQuery.Message.Chat.ID)
		}
		return b.startCreatingTask(callbackQuery, callback)
	case entity.UpdateTask:
		if !b.isAdmin(callbackQuery.From.UserName) {
			return b.notAdminResponse(callbackQuery.Message.Chat.ID)
		}
		return b.startUpdateTask(callbackQuery, callback)
	case entity.DeleteTask:
		if !b.isAdmin(callbackQuery.From.UserName) {
			return b.notAdminResponse(callbackQuery.Message.Chat.ID)
		}
		return b.deleteTask(callbackQuery, callback)
	case entity.UpdateSubject:
		if !b.isAdmin(callbackQuery.From.UserName) {
			return b.notAdminResponse(callbackQuery.Message.Chat.ID)
		}
		return b.startUpdateSubject(callbackQuery, callback)
	case entity.DeleteSubject:
		if !b.isAdmin(callbackQuery.From.UserName) {
			return b.notAdminResponse(callbackQuery.Message.Chat.ID)
		}
		return b.deleteSubject(callbackQuery, callback)
	}
	return nil
}
