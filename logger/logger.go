package logger

import (
	"github.com/sirupsen/logrus"
)

const (
	loggerPrefix = "_logger"
)

type hook struct {
	logger string
}

// New creates a hook that will add a value to the loggerPrefix field
// Multiple New calls will just change the value if it already exists
func New(logger string) logrus.Hook {
	return &hook{logger: logger}
}

//AddLogger just adds a `_logger` with field.
func AddLogger(logger *logrus.Logger, loggervalue string) {
	logger.Hooks.Add(New(loggervalue))
}

func (h *hook) Fire(entry *logrus.Entry) error {
	entry.Data[loggerPrefix] = h.logger
	return nil
}

func (h *hook) Levels() []logrus.Level {
	return logrus.AllLevels
}
