package main

import (
	"context"
	"fmt"
	"github.com/patriarch11/telegram-task-manager-bot/internal/config"
	"github.com/patriarch11/telegram-task-manager-bot/internal/di"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/ypopivniak/envconfig"
	"path/filepath"
	"runtime"
)

func main() {
	var conf config.Config
	var err error

	ctx := context.Background()

	err = envconfig.Process("", &conf)
	if err != nil {
		logrus.Fatal(err)
	}

	level, err := logrus.ParseLevel(conf.LogLevel)
	if err != nil {
		logrus.Fatal(errors.Wrap(err, "invalid log level - should be one of fatal, error, warn, info, debug, trace"))
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d\t", filepath.Base(f.File), f.Line)
		},
	})

	/*sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)*/
	// TODO: add close func
	bot, err := di.InitializeBot(ctx, &conf)
	if err != nil {
		logrus.Fatal(err)
	}
	bot.HandleUpdates()

}
