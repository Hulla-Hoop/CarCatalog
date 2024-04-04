package carcatalog

import (
	"carcatalog/internal/model"
	"encoding/json"
	"io"
	"net/http"
)

func (c *carcatalog) InsertCar(reqId string, regNum []string) error {
	c.logger.WithField("carCatalog.InsertCar", reqId).Debug("полученные данные ", regNum)

	for _, reg := range regNum {
		car, err := c.get(reqId, reg)
		if err != nil {
			return err
		}
		c.db.Insert(reqId, *car)
	}

	return nil
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
