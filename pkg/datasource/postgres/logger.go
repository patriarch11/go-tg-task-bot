package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	l logrus.FieldLogger
}

func NewLogger(l logrus.FieldLogger) *Logger {
	return &Logger{l: l}
}

func (l *Logger) Log(_ context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	var logger logrus.FieldLogger
	if data != nil {
		logger = l.l.WithFields(data)
	} else {
		logger = l.l
	}

	switch level {
	case pgx.LogLevelTrace:
		logger.WithField("PGX_LOG_LEVEL", level).Debug(msg)
	case pgx.LogLevelDebug, pgx.LogLevelInfo:
		logger.Debug(msg)
	case pgx.LogLevelWarn:
		logger.Warn(msg)
	case pgx.LogLevelError:
		logger.Error(msg)
	default:
		logger.WithField("INVALID_PGX_LOG_LEVEL", level).Error(msg)
	}
}
