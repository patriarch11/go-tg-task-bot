package telegram

import (
	"github.com/gofrs/uuid"
)

func (b *Bot) showTasks(chatID int64, subjectID uuid.UUID) error {
	tasks, err := b.taskRepository.GetAllBySubjectID(subjectID)
	if err != nil {
		return err
	}
	for _, task := range tasks {
		msg := taskMessage(chatID, *task)
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
	}
	return nil
}
