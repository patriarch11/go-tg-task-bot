package task

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/go-tg-task-bot/internal/domain/entity"
	"github.com/patriarch11/go-tg-task-bot/internal/interfaces"
)

type UseCaseTask struct {
	task            *entity.Task
	taskService     interfaces.TaskService
	keyboardService interfaces.KeyboardService
}

func NewUseCaseTask(
	taskService interfaces.TaskService,
	keyboardService interfaces.KeyboardService) *UseCaseTask {
	return &UseCaseTask{
		task:            &entity.Task{},
		taskService:     taskService,
		keyboardService: keyboardService,
	}
}

func (u *UseCaseTask) ShowTasks(handler interfaces.UpdateHandler,
	subjectId entity.ID, chatId int64, isAdmin bool) error {
	tasks, err := u.taskService.GetBySubjectId(context.Background(), subjectId)
	if err != nil {
		return err
	}
	for _, task := range tasks {
		rep := tgbotapi.NewMessage(chatId, task.Description)
		rep = u.keyboardService.WrapTaskMessageInInlineKeyboard(rep, task, isAdmin)
		_, err = handler.BotAPI().Send(rep)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *UseCaseTask) AddTaskReply(handler interfaces.UpdateHandler, subjectId entity.ID, chatId int64) error {
	u.task.SubjectId = subjectId
	rep := tgbotapi.NewMessage(chatId, entity.InputTaskDescriptionReply)
	rep = u.keyboardService.WrapMessageInCancelKeyboard(rep)

	handler.SetState(entity.ReceiveTask)

	_, err := handler.BotAPI().Send(rep)
	if err != nil {
		u.task = &entity.Task{}
		handler.SetState(entity.Default)
		return err
	}
	return nil
}

func (u *UseCaseTask) ReceiveTaskDescriptionAndSave(handler interfaces.UpdateHandler, msg *tgbotapi.Message) error {

	u.task.Description = msg.Text
	handler.SetState(entity.Default)

	_, err := u.taskService.Create(context.Background(), u.task)
	if err != nil {
		u.task = &entity.Task{}
		handler.SetState(entity.Default)
		return err
	}

	rep := tgbotapi.NewMessage(msg.Chat.ID, entity.TaskAddedReply)
	rep = u.keyboardService.WrapMessageInMainKeyboard(rep, true)

	_, err = handler.BotAPI().Send(rep)
	if err != nil {
		u.task = &entity.Task{}
		handler.SetState(entity.Default)
		return err
	}
	u.task = &entity.Task{}
	return nil
}
func (u *UseCaseTask) UpdateTaskReply(handler interfaces.UpdateHandler, taskId entity.ID, chatId int64) error {
	u.task.Id = taskId
	rep := tgbotapi.NewMessage(chatId, entity.InputTaskDescriptionReply)
	rep = u.keyboardService.WrapMessageInCancelKeyboard(rep)
	handler.SetState(entity.ReceiveUpdTask)
	_, err := handler.BotAPI().Send(rep)
	if err != nil {
		handler.SetState(entity.Default)
		return err
	}
	return nil
}

func (u *UseCaseTask) ReceiveUpdTaskDescriptionAndSave(handler interfaces.UpdateHandler, msg *tgbotapi.Message) error {

	u.task.Description = msg.Text
	handler.SetState(entity.Default)

	_, err := u.taskService.Update(context.Background(), u.task)
	if err != nil {
		u.task = &entity.Task{}
		return err
	}

	rep := tgbotapi.NewMessage(msg.Chat.ID, entity.TaskUpdatedReply)
	rep = u.keyboardService.WrapMessageInMainKeyboard(rep, true)

	_, err = handler.BotAPI().Send(rep)
	if err != nil {
		u.task = &entity.Task{}
		handler.SetState(entity.Default)
		return err
	}
	u.task = &entity.Task{}
	return nil
}

func (u *UseCaseTask) DeleteTask(handler interfaces.UpdateHandler, taskId entity.ID, chatId int64) error {
	err := u.taskService.Delete(context.Background(), taskId)
	if err != nil {
		return err
	}
	rep := tgbotapi.NewMessage(chatId, entity.TaskDeletedReply)
	rep = u.keyboardService.WrapMessageInMainKeyboard(rep, true)
	_, err = handler.BotAPI().Send(rep)
	if err != nil {
		return err
	}
	return nil
}
