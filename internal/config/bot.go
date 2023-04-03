package config

import "github.com/patriarch11/go-tg-task-bot/pkg/datasource/postgres"

type Config struct {
	LogLevel string           `required:"true" default:"info" split_words:"true"`
	Token    string           `required:"true"`
	Admin    string           `required:"true"`
	Database *postgres.Config `required:"true" split_words:"true"`
}
