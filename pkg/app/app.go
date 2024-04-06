package app

import (
	"carcatalog/internal/REPO/psql"
	"carcatalog/internal/config"
	cataloge "carcatalog/internal/endpoint/catalogE"
	"carcatalog/internal/logger"
	"carcatalog/internal/service/carcatalog"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type app struct {
	mux *http.ServeMux
	l   *logrus.Logger
}

func New() *app {
	l := logger.New()

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	l.WithField("app.New", "start").Info("Конфигурации загружены")
	db, err := psql.InitDb(l)
	if err != nil {
		panic(err)
	}
	l.WithField("app.New", "start").Info("Бд поднялась")

	s := carcatalog.InitCarCatalog(l, db)
	l.WithField("app.New", "start").Info("сервисный слой поднялся")

	h := cataloge.Init(l, s)

	mux := http.NewServeMux()

	mux.HandleFunc("/insert", h.Insert)
	mux.HandleFunc("/update", h.Update)
	mux.HandleFunc("/delete", h.Delete)
	mux.HandleFunc("/filter", h.Filter)

	l.WithField("app.New", "start").Info("эндпоинты готовы")

	return &app{
		mux: mux,
		l:   l,
	}

}

func (a *app) Start() {
	conf := config.ServNew()
	a.l.WithField("APP.Start", "").Infof("Сервер стартовал на %s:%s", conf.Host, conf.Port)
	err := http.ListenAndServe(conf.Host+":"+conf.Port, a.mux)
	if err != nil {
		a.l.WithField("APP.Start", err).Error()
	}

}
