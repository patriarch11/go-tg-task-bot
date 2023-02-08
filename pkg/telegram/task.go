package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofrs/uuid"
	"github.com/patriarch11/go-tg-task-bot/internal/entity"
	"log"
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

func (b *Bot) startCreatingTask(callbackQuery *tgbotapi.CallbackQuery, callback entity.Callback) error {
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Введіть опис завдання")
	b.state = WaitForTaskDescription
	b.task.SubjectId = callback.ID
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) startUpdateTask(callbackQuery *tgbotapi.CallbackQuery, callback entity.Callback) error {
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Введіть опис завдання")
	b.state = WaitForNewTaskDescription
	b.task.ID = callback.ID
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) updateTaskDescription(message *tgbotapi.Message) error {
	b.task.Description = message.Text
	task, err := b.taskRepository.Update(&b.task)
	if err != nil {
		log.Println(err)
		return err
	}
	msg := taskMessage(message.Chat.ID, *task)
	b.state = None
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (b *Bot) createTask(message *tgbotapi.Message) error {
	b.task.Description = message.Text
	task, err := b.taskRepository.Create(&b.task)
	if err != nil {
		log.Println(err)
		return err
	}
	msg := taskMessage(message.Chat.ID, *task)
	b.state = None
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil

}

func (b *Bot) deleteTask(callbackQuery *tgbotapi.CallbackQuery, callback entity.Callback) error {
	err := b.taskRepository.Delete(callback.ID)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "Завдання видалено")
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}
