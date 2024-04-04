package logger

import (
	formatter "github.com/fabienm/go-logrus-formatters"
	"github.com/sirupsen/logrus"
)

func New() *logrus.Logger {
	logger := logrus.New()
	rr := formatter.NewGelf("Catalog")
	logger.SetFormatter(rr)
	logger.SetLevel(logrus.DebugLevel)
	return logger
}
