package task

import "github.com/patriarch11/telegram-task-manager-bot/internal/interfaces"

type UseCaseTask struct {
	subjectService interfaces.SubjectService
	taskService    interfaces.TaskService
}
