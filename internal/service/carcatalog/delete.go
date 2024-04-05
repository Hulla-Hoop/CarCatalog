package carcatalog

import (
	"carcatalog/internal/model"
	"errors"
	"regexp"
	"strconv"
)

func (c *carcatalog) Delete(reqId string, id string) (*model.Car, error) {

	idi, err := c.idCheckAndConvert(reqId, id)
	if err != nil {
		return nil, err
	}
	carDb, err := c.db.Delete(reqId, idi)
	if err != nil {
		return nil, err
	}
	car := carDb.CarToCar()

	return car, nil

}

func (c *carcatalog) idCheckAndConvert(reqId string, id string) (int, error) {
	pattern := `^[0-9]*$`
	match, _ := regexp.MatchString(pattern, id)
	if !match {
		c.logger.WithField("carCatalog.idCheckAndConvert", reqId).Error("некорректные данные ", id)
		return 0, errors.New("некорректные данные")
	}
	idi, err := strconv.Atoi(id)
	if err != nil {
		c.logger.WithField("carCatalog.idCheckAndConvert", reqId).Error("Ннеудалось преобразовать ", id)
		return 0, err
	}

	return idi, nil
}
