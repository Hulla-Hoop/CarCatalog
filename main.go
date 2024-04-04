package main

import (
	"carcatalog/internal/REPO/psql"
	cataloge "carcatalog/internal/endpoint/catalogE"
	"carcatalog/internal/logger"
	"carcatalog/internal/service/carcatalog"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	log := logger.New()
	db, err := psql.InitDb(log)
	if err != nil {
		log.Fatal(err)
	}
	carcatalog := carcatalog.InitCarCatalog(log, db)

	endpoint := cataloge.Init(log, carcatalog)

	http.HandleFunc("/insert", endpoint.Insert)

	err = http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}

}
