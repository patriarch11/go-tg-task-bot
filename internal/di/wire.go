//go:build wireinject
// +build wireinject

package di

import (
	"context"
	"github.com/google/wire"
	"github.com/jackc/pgx/v4/pgxpool"
	storage "github.com/patriarch11/go-tg-task-bot/internal/adapters/db/postgres"
	"github.com/patriarch11/go-tg-task-bot/internal/config"
	"github.com/patriarch11/go-tg-task-bot/internal/controller/telegram"
	"github.com/patriarch11/go-tg-task-bot/internal/domain/service"
	"github.com/patriarch11/go-tg-task-bot/internal/domain/usecase/subject"
	"github.com/patriarch11/go-tg-task-bot/internal/domain/usecase/task"
	"github.com/patriarch11/go-tg-task-bot/internal/interfaces"
	"github.com/patriarch11/go-tg-task-bot/pkg/datasource/postgres"
	"github.com/ypopivniak/pgxevents/pkg/pgxevents"
)

var DatasourceSet = wire.NewSet(
	postgres.New,
	pgxevents.NewListener,
	wire.Value([]pgxevents.Option{}),
	wire.FieldsOf(new(*config.Config), "Database"),
	wire.Bind(new(postgres.Datasource), new(*pgxpool.Pool)),
)

var SubjectSet = wire.NewSet(
	storage.NewSubjectStorage,
	service.NewSubjectService,
	service.NewKeyboardService,
	subject.NewUseCaseSubject,
	wire.Bind(new(interfaces.SubjectStorage), new(*storage.SubjectStorage)),
	wire.Bind(new(interfaces.SubjectService), new(*service.SubjectService)),
	wire.Bind(new(interfaces.KeyboardService), new(*service.KeyboardService)),
	wire.Bind(new(interfaces.UseCaseSubject), new(*subject.UseCaseSubject)),
)

var TaskSet = wire.NewSet(
	storage.NewTaskStorage,
	service.NewTaskService,
	task.NewUseCaseTask,
	wire.Bind(new(interfaces.TaskStorage), new(*storage.TaskStorage)),
	wire.Bind(new(interfaces.TaskService), new(*service.TaskService)),
	wire.Bind(new(interfaces.UseCaseTask), new(*task.UseCaseTask)),
)

var BotSet = wire.NewSet(
	telegram.NewUpdateHandler,
	wire.Bind(new(interfaces.UpdateHandler), new(*telegram.UpdateHandler)),
)

func InitializeBot(ctx context.Context, conf *config.Config) (*telegram.UpdateHandler, error) {
	wire.Build(
		DatasourceSet,
		SubjectSet,
		TaskSet,
		BotSet,
	)
	return new(telegram.UpdateHandler), nil
}
