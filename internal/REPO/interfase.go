package repo

import "carcatalog/internal/model"

type Repo interface {
	Insert(reqId string, car model.Car) (*model.CarDB, error)
	Delete(reqId string, id int) (*model.CarDB, error)
	Update(reqId string, car model.Car) (*model.CarDB, error)
	Filter(reqId string, filter map[string]string) ([]model.CarDB, error)
}
