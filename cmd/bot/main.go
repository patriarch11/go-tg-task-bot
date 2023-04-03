package main

import (
	"github.com/patriarch11/telegram-task-manager-bot/internal/config"
	"github.com/patriarch11/telegram-task-manager-bot/internal/controller/telegram"
	"github.com/patriarch11/telegram-task-manager-bot/internal/domain/service"
	"github.com/sirupsen/logrus"
	"github.com/ypopivniak/envconfig"
)

func main() {
	var conf config.BotConfig
	var err error

	err = envconfig.Process("", &conf)
	if err != nil {
		logrus.Fatal(err)
	}
	messageService := service.NewMessageService(conf)

	updateHandler, err := telegram.NewUpdateHandler(conf, messageService)
	if err != nil {
		logrus.Fatal(err)
	}
	updateHandler.HandleUpdates()
}
