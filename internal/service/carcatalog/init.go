package carcatalog

import (
	repo "carcatalog/internal/REPO"
	"carcatalog/internal/config"

	"github.com/sirupsen/logrus"
)

type carcatalog struct {
	logger *logrus.Logger
	db     repo.Repo
	cfg    *config.ConfigRemoteApi
}

func InitCarCatalog(logger *logrus.Logger, db repo.Repo) *carcatalog {
	cfg := config.RemoteApi()
	return &carcatalog{
		logger: logger,
		db:     db,
		cfg:    cfg,
	}
}
