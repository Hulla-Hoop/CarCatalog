package cataloge

import (
	"carcatalog/internal/service"

	"github.com/sirupsen/logrus"
)

type endpoint struct {
	s      service.Service
	logger *logrus.Logger
}

func Init(logger *logrus.Logger, s service.Service) *endpoint {
	return &endpoint{
		s:      s,
		logger: logger,
	}
}
