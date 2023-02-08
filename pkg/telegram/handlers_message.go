package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	switch b.state {
	case WaitForSubjectName:
		return b.startCreatingSubject(message)
	case WaitForNewSubjectName:
		return b.updateSubjectName(message)
	case WaitForNewSubjectDescription:
		return b.updateSubject(message)
	case WaitForSubjectDescription:
		return b.createSubject(message)
	case WaitForTaskDescription:
		return b.createTask(message)
	case WaitForNewTaskDescription:
		return b.updateTaskDescription(message)
	default:
		return nil
	}
}
