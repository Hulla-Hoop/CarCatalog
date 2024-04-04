package service

import (
	"carcatalog/internal/model"
)

type Service interface {
	Insert(reqId string, regNum []string) ([]model.CarDB, error)
}
