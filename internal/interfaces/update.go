package interfaces

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/telegram-task-manager-bot/internal/domain/entity"
)

type UpdateHandler interface {
	HandleUpdates()
	SetState(state entity.State)
	State() entity.State
	BotAPI() *tgbotapi.BotAPI
}
