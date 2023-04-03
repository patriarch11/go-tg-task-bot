package config

type BotConfig struct {
	Token string `required:"true"`
	Admin string `required:"true"`
}
