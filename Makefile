ifneq ($(wildcard .env.example),)
    include .env.example
endif

export

.PHONY: run
run: ## Run bot
	go run github.com/patriarch11/go-tg-task-bot/cmd/bot

.PHONY: migrate-up
migrate-up: ## Run migration up
	migrate -verbose -database "${DATABASE_URL}" -path migrations up
