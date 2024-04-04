package service

import (
	"carcatalog/internal/model"
)

type Service interface {
	Insert(reqId string, car model.Car) (*model.Car, error)
}
