package carcatalog

import (
	"carcatalog/internal/model"
	"regexp"
)

func (c *carcatalog) Filter(reqId string, limit string, offset string, field string, value string, operator string) ([]model.Car, error) {

	operators := c.checkFilter(reqId, limit, offset, field, value, operator)

	carsDB, err := c.db.Filter(reqId, operators)

	if err != nil {
		return nil, err
	}

	var cars []model.Car

	for _, carDB := range carsDB {
		car := carDB.CarToCar()
		cars = append(cars, *car)
	}

	return cars, nil
}

func (c *carcatalog) checkFilter(reqId string, limit string, offset string, field string, value string, operator string) map[string]string {

	c.logger.WithField("carCatalog.checkFilter", reqId).Debug("Полученные данные Limit: ", limit, " Offset: ", offset, " Field: ", field, " Value: ", value, " Operator: ", operator)

	operators := make(map[string]string)
	if limit != "" {
		match, err := regexp.MatchString(`^[0-9]*$`, limit)
		if err != nil {
			c.logger.WithField("carCatalog.checkFilter", reqId).Error("некорректные данные ", limit)
		} else if !match {
			c.logger.WithField("carCatalog.checkFilter", reqId).Error("некорректные данные ", limit)
		} else {
			operators["limit"] = limit
		}
	}
	if offset != "" {
		match, err := regexp.MatchString(`^[0-9]*$`, offset)
		if err != nil {
			c.logger.WithField("carCatalog.checkFilter", reqId).Error("некорректные данные ", offset)
		} else if !match {
			c.logger.WithField("carCatalog.checkFilter", reqId).Error("некорректные данные ", offset)
		} else {
			operators["offset"] = offset
		}
	}
	if field == "regNum" || field == "mark" || field == "model" || field == "year" || field == "name" || field == "surname" || field == "patronymic" {
		operators["field"] = field
	}
	if value != "" {
		operators["value"] = value
	}
	if operator == "eq" || operator == "ne" || operator == "gt" || operator == "ge" || operator == "lt" || operator == "le" {
		operators["operator"] = operator
	}

	return operators

}
