package repo

import "carcatalog/internal/model"

type Repo interface {
	Insert(reqId string, car model.Car) (*model.CarDB, error)
}
