package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/go-tg-task-bot/internal/entity"
	"github.com/patriarch11/go-tg-task-bot/internal/protocol"
	"github.com/patriarch11/go-tg-task-bot/internal/repository"
	"github.com/patriarch11/go-tg-task-bot/pkg/datasource"
	"log"
)

type State int

const (
	None State = iota
	WaitForSubjectName
	WaitForNewSubjectName
	WaitForSubjectDescription
	WaitForNewSubjectDescription
	WaitForTaskDescription
	WaitForNewTaskDescription
)

type Bot struct {
	bot               *tgbotapi.BotAPI
	taskRepository    protocol.PostgresTaskRepository
	subjectRepository protocol.PostgresSubjectRepository
	adminUserName     string
	state             State
	subject           entity.Subject
	task              entity.Task
}

func NewBot(bot *tgbotapi.BotAPI, userName string,
	datasource *datasource.Datasource) *Bot {
	taskRepository := repository.NewPostgresTaskRepository(datasource)
	subjectRepository := repository.NewPostgresSubjectRepository(datasource)
	return &Bot{
		bot:               bot,
		adminUserName:     userName,
		taskRepository:    taskRepository,
		subjectRepository: subjectRepository}
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)
	b.handleUpdates(b.initUpdatesChannel())
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	// TODO: must handle errors
	for update := range updates {
		if update.CallbackQuery != nil {
			_ = b.handleCallback(update.CallbackQuery)
			continue
		}
		if update.Message.IsCommand() {
			_ = b.handleCommand(update.Message)
			continue
		}
		if update.Message.Text != "" {
			_ = b.handleMessage(update.Message)
		}
	}
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	return updates
}

func (b *Bot) isAdmin(userName string) bool {
	return b.adminUserName == userName
}

func (b *Bot) notAdminResponse(chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "permission denied")
	_, err := b.bot.Send(msg)
	return err
}
