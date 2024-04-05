package service

import (
	"carcatalog/internal/model"
)

type Service interface {
	Insert(reqId string, regNum []string) ([]model.Car, error)
	Delete(reqId string, id string) (*model.Car, error)
	Update(reqId string, car model.Car, id string) (*model.Car, error)
	Filter(reqId string, limit string, offset string, field string, value string, operator string) ([]model.Car, error)
}
