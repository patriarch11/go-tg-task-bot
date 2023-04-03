package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/go-tg-task-bot/internal/config"
	"github.com/patriarch11/go-tg-task-bot/internal/controller/dto"
	"github.com/patriarch11/go-tg-task-bot/internal/domain/entity"
	"github.com/patriarch11/go-tg-task-bot/internal/interfaces"
	"github.com/sirupsen/logrus"
)

const (
	startCommand = "start"
	helpCommand  = "help"
)

type UpdateHandler struct {
	state           entity.State
	adminUsername   string
	bot             *tgbotapi.BotAPI
	updates         tgbotapi.UpdatesChannel
	keyboardService interfaces.KeyboardService
	subjectManager  interfaces.UseCaseSubject
	taskManager     interfaces.UseCaseTask
}

func NewUpdateHandler(config *config.Config,
	keyboardService interfaces.KeyboardService,
	subjectManager interfaces.UseCaseSubject,
	taskManager interfaces.UseCaseTask,

) (updateHandler *UpdateHandler, err error) {

	updateHandler = new(UpdateHandler)
	updateHandler.adminUsername = config.Admin
	updateHandler.keyboardService = keyboardService
	updateHandler.subjectManager = subjectManager
	updateHandler.taskManager = taskManager
	updateHandler.state = entity.Default

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
		if update.CallbackQuery == nil && update.Message == nil {
			continue
		}
		if update.CallbackQuery != nil {
			err = h.handleCallback(update.CallbackQuery)
			if err != nil {
				logrus.Error(err)
			}
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
			err = h.handleTextMessage(update.Message)
			if err != nil {
				logrus.Error(err)
			}
			continue
		}
	}
}

func (h *UpdateHandler) SetState(state entity.State) {
	h.state = state
}

func (h *UpdateHandler) State() entity.State {
	return h.state
}

func (h *UpdateHandler) BotAPI() *tgbotapi.BotAPI {
	return h.bot
}

func (h *UpdateHandler) handleCommand(commandMessage *tgbotapi.Message) error {
	switch commandMessage.Command() {
	case startCommand:
		err := h.replyWithMainKeyboard(commandMessage.Chat.ID, entity.StartCommandReply,
			h.isAdmin(commandMessage.From.UserName))
		if err != nil {
			return err
		}
		return nil
	case helpCommand:
		err := h.replyWithMainKeyboard(commandMessage.Chat.ID, entity.HelpCommandReply,
			h.isAdmin(commandMessage.From.UserName))
		if err != nil {
			return err
		}
		return nil
	default:
		err := h.replyWithMainKeyboard(commandMessage.Chat.ID, entity.UnknownCommandReply,
			h.isAdmin(commandMessage.From.UserName))
		if err != nil {
			return err
		}
		return nil
	}
}

func (h *UpdateHandler) handleTextMessage(textMessage *tgbotapi.Message) error {
	username := textMessage.From.UserName

	if textMessage.Text == entity.CancelButtonText && h.isAdmin(username) {
		h.SetState(entity.Default)
		return h.replyWithMainKeyboard(textMessage.Chat.ID, entity.CanceledReply, true)
	}

	switch h.State() {
	case entity.Default:

		switch textMessage.Text {
		case entity.AddSubjectButtonText:
			if !h.isAdmin(username) {
				return h.replyWithMainKeyboard(textMessage.Chat.ID, entity.PermissionDeniedReply, false)
			}
			return h.subjectManager.AddSubjectReply(h, textMessage)
		case entity.ShowSubjectsButtonText:
			return h.subjectManager.ShowAllSubjects(h, textMessage, h.isAdmin(username))
		default:
			return h.replyWithMainKeyboard(textMessage.Chat.ID, textMessage.Text, h.isAdmin(username))
		}

	case entity.ReceiveSubjectName:
		return h.subjectManager.SetSubjectName(h, textMessage)
	case entity.ReceiveSubjectDescription:
		return h.subjectManager.ReceiveSubjectDescriptionAndSave(h, textMessage)
	case entity.ReceiveTask:
		return h.taskManager.ReceiveTaskDescriptionAndSave(h, textMessage)
	case entity.ReceiveUpdSubjectName:
		return h.subjectManager.SetUpdSubjectName(h, textMessage)
	case entity.ReceiveUpdSubjectDescription:
		return h.subjectManager.ReceiveUpdSubjectDescriptionAndSave(h, textMessage)
	case entity.ReceiveUpdTask:
		return h.taskManager.ReceiveUpdTaskDescriptionAndSave(h, textMessage)

	}
	return nil
}

func (h *UpdateHandler) handleCallback(callbackQuery *tgbotapi.CallbackQuery) error {
	c := tgbotapi.NewCallback(callbackQuery.ID, callbackQuery.Data)
	if _, err := h.bot.Request(c); err != nil {
		return err
	}
	callback := dto.CallbackFromString(callbackQuery.Data)

	chatId := callbackQuery.Message.Chat.ID
	username := callbackQuery.From.UserName

	switch callback.OperationType {
	case entity.ShowTasks:
		return h.taskManager.ShowTasks(h, callback.Id, chatId, h.isAdmin(username))
	case entity.AddTask:
		return h.taskManager.AddTaskReply(h, callback.Id, chatId)
	case entity.UpdateSubject:
		return h.subjectManager.UpdateSubjectReply(h, callback.Id, chatId)
	case entity.DeleteSubject:
		return h.subjectManager.DeleteSubject(h, callback.Id, chatId)
	case entity.UpdateTask:
		return h.taskManager.UpdateTaskReply(h, callback.Id, chatId)
	case entity.DeleteTask:
		return h.taskManager.DeleteTask(h, callback.Id, chatId)
	default:
		return nil
	}
}

func (h *UpdateHandler) replyWithMainKeyboard(senderChatId int64, reply string, isAdmin bool) error {
	msg := tgbotapi.NewMessage(senderChatId, reply)
	msg = h.keyboardService.WrapMessageInMainKeyboard(msg, isAdmin)
	_, err := h.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (h *UpdateHandler) isAdmin(username string) bool {
	return h.adminUsername == username
}
