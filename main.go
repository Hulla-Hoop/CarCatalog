package main

import (
	"carcatalog/internal/REPO/psql"
	"carcatalog/internal/logger"
	"carcatalog/internal/service/carcatalog"

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

	sl := []string{"A111AA111", "A222AA222", "A333AA333", "A444AA444"}

	carcatalog.InsertCar("1", sl)
}
