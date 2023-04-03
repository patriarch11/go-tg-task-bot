package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/telegram-task-manager-bot/internal/config"
	"github.com/patriarch11/telegram-task-manager-bot/internal/interfaces"
	"github.com/sirupsen/logrus"
)

const (
	startCommand = "start"
	helpCommand  = "help"
)

type UpdateHandler struct {
	bot            *tgbotapi.BotAPI
	updates        tgbotapi.UpdatesChannel
	messageService interfaces.MessageService
}

func NewUpdateHandler(config config.BotConfig, messageService interfaces.MessageService) (updateHandler *UpdateHandler, err error) {
	updateHandler = new(UpdateHandler)
	updateHandler.messageService = messageService
	updateHandler.bot, err = tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		return nil, err
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updateHandler.updates = updateHandler.bot.GetUpdatesChan(updateConfig)
	return
}

func (h *UpdateHandler) HandleUpdates() {
	var err error

	for update := range h.updates {
		if update.CallbackQuery != nil {
			h.handleCallback(update.CallbackQuery)
			continue
		}
		if update.Message.IsCommand() {
			err = h.handleCommand(update.Message)
			if err != nil {
				logrus.Error(err)
			}
			continue
		}
		if update.Message.Text != "" {
			h.handleTextMessage(update.Message)
			continue
		}
	}
}

func (h *UpdateHandler) handleCommand(commandMessage *tgbotapi.Message) error {
	switch commandMessage.Command() {
	case startCommand:
		err := h.replyToCommand(commandMessage.Chat.ID, startCommandReply, commandMessage.From.UserName)
		if err != nil {
			return err
		}
		return nil
	case helpCommand:
		err := h.replyToCommand(commandMessage.Chat.ID, helpCommandReply, commandMessage.From.UserName)
		if err != nil {
			return err
		}
		return nil
	default:
		err := h.replyToCommand(commandMessage.Chat.ID, unknownCommandReply, commandMessage.From.UserName)
		if err != nil {
			return err
		}
		return nil
	}
}

func (h *UpdateHandler) handleTextMessage(textMessage *tgbotapi.Message) {

}

func (h *UpdateHandler) handleCallback(callback *tgbotapi.CallbackQuery) {

}

func (h *UpdateHandler) replyToCommand(senderChatId int64, reply, username string) error {
	msg := tgbotapi.NewMessage(senderChatId, reply)
	msg = h.messageService.WrapMessageInMainKeyboard(msg, username)
	_, err := h.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}
