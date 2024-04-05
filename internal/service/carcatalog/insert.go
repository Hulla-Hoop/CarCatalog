package carcatalog

import (
	"carcatalog/internal/model"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
)

func (c *carcatalog) Insert(reqId string, regNum []string) ([]model.Car, error) {
	c.logger.WithField("carCatalog.InsertCar", reqId).Debug("полученные данные ", regNum)
	var cars []model.Car
	for _, reg := range regNum {

		if !c.check(reg) {
			c.logger.WithField("carCatalog.InsertCar", reqId).Debug("некорректные данные ", reg)
			continue
		}

		car, err := c.get(reqId, reg)
		if err != nil {
			return nil, err
		}
		carDB, err := c.db.Insert(reqId, *car)
		if err != nil {
			return nil, err
		}

		cars = append(cars, *carDB.CarToCar())
	}

	return cars, nil
}

func (c *carcatalog) get(reqId string, regNum string) (*model.Car, error) {

	var car model.Car

	r, err := http.Get(c.cfg.Link + regNum)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &car)
	c.logger.WithField("carCatalog.get", reqId).Debug(car)
	return &car, nil
}

func (c *carcatalog) check(regNum string) bool {
	pattern := `^[a-zA-Zа-яА-Я]\d{3}[a-zA-Zа-яА-Я][a-zA-Zа-яА-Я]\d{3}$`
	match, _ := regexp.MatchString(pattern, regNum)
	return match
}
