package main

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/patriarch11/go-tg-task-bot/pkg/datasource"
	"github.com/patriarch11/go-tg-task-bot/pkg/telegram"
	"log"
	"os"
)

//set -a && source .env && make run

func main() {
	ctx := context.Background()
	db, err := datasource.New(ctx, mustDBConfig())
	if err != nil {
		log.Fatal(err)
	}
	bot, err := tgbotapi.NewBotAPI(mustToken())
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	b := telegram.NewBot(bot, mustAdminUserName(), db)
	b.Start()
}

func mustToken() string {
	token := os.Getenv("API_BOT_TOKEN")
	if token == "" {
		log.Fatal("API bot token is required")
	}
	return token
}

func mustAdminUserName() string {
	name := os.Getenv("ADMIN_USERNAME")
	if name == "" {
		log.Fatal("Admin username is required")
	}
	return name
}

func mustDBConfig() *datasource.Config {
	url := os.Getenv("DB_URL")
	if url == "" {
		log.Fatal("DB URL is required")
	}
	return &datasource.Config{
		URL:                  url,
		MaxConnections:       10,
		IdleConnections:      5,
		PreferSimpleProtocol: true,
	}
}
