package dto

import (
	"encoding/json"
	"github.com/patriarch11/go-tg-task-bot/internal/domain/entity"
	"github.com/sirupsen/logrus"
)

func CallbackFromString(str string) *entity.Callback {
	var data *entity.Callback
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		logrus.Error("unmarshaling error: %s", err)
	}
	return data
}
