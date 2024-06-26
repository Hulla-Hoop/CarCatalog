package psql

import (
	"carcatalog/internal/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
)

type psql struct {
	dB     *sql.DB
	logger *logrus.Logger
}

func InitDb(logger *logrus.Logger) (*psql, error) {
	config := config.DbNew()

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=%s", config.Host, config.User, config.DBName, config.Password, config.Port, config.SSLMode)

	dB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = goose.Up(dB, "migrations")
	if err != nil {
		return nil,
			fmt.Errorf("--- Ошибка миграции:%s", err)
	}
	return &psql{
		dB:     dB,
		logger: logger,
	}, nil

}
