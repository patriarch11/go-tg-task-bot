package dto

import (
	"encoding/json"
	"github.com/patriarch11/go-tg-task-bot/internal/entity"
	"log"
)

func CallbackFromString(str string) entity.Callback {
	var data entity.Callback
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		log.Printf("unmarshaling error: %s", err)
	}
	return data
}
