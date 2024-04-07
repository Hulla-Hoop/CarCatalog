package carcatalog

import (
	"carcatalog/internal/model"
	"errors"
	"regexp"
)

func (c *carcatalog) Update(reqId string, car model.Car, id string) (*model.Car, error) {

	idi, err := c.idCheckAndConvert(reqId, id)

	if err != nil {
		return nil, err
	}

	car.Id = idi

	err = c.checkcar(reqId, &car)
	if err != nil {
		return nil, err
	}

	carDB, err := c.db.Update(reqId, car)
	if err != nil {
		return nil, err
	}

	car = *carDB.CarToCar()

	return &car, nil
}

// проверяет валидность полей структуры Car в случае несоответствия возвращает ошибку
func (c *carcatalog) checkcar(reqId string, car *model.Car) error {
	if car.RegNum != "" {
		match, err := regexp.MatchString(`^[a-zA-Zа-яА-Я]\d{3}[a-zA-Zа-яА-Я][a-zA-Zа-яА-Я]\d{3}$`, car.RegNum)
		if err != nil {
			return err
		}
		if !match {
			c.logger.WithField("carCatalog.checkcar", reqId).Debug("некорректные данные RegNum ", car.RegNum)
			return errors.New("некорректные данные RegNum")
		}
	}
	if car.Model != "" {
		match, err := regexp.MatchString(`^[a-zA-Zа-яА-Я0-9]`, car.Model)
		if err != nil {
			return err
		}
		if !match {
			c.logger.WithField("carCatalog.checkcar", reqId).Debug("некорректные данные Model ", car.Model)
			return errors.New("некорректные данные Model")
		}
	}
	if car.Mark != "" {
		match, err := regexp.MatchString(`^[a-zA-Zа-яА-Я0-9]`, car.Mark)
		if err != nil {
			return err
		}
		if !match {
			c.logger.WithField("carCatalog.checkcar", reqId).Debug("некорректные данные Mark ", car.Mark)
			return errors.New("некорректные данные Mark")
		}
	}
	if car.Owner == nil {
		car.Owner = &model.People{}
		return nil
	} else {
		if car.Owner.Name != "" {
			match, err := regexp.MatchString(`^[a-zA-Zа-яА-Я]`, car.Owner.Name)
			if err != nil {
				return err
			}
			if !match {
				c.logger.WithField("carCatalog.checkcar", reqId).Debug("некорректные данные Name ", car.Owner.Name)
				return errors.New("некорректные данные Name")
			}
		}
		if car.Owner.Surname != "" {
			match, err := regexp.MatchString(`^[a-zA-Zа-яА-Я]`, car.Owner.Surname)
			if err != nil {
				return err
			}
			if !match {
				c.logger.WithField("carCatalog.checkcar", reqId).Debug("некорректные данные Surname ", car.Owner.Surname)
				return errors.New("некорректные данные Surname")
			}
		}
		if car.Owner.Patronymic != "" {
			match, err := regexp.MatchString(`^[a-zA-Zа-яА-Я]`, car.Owner.Patronymic)
			if err != nil {
				return err
			}
			if !match {
				c.logger.WithField("carCatalog.checkcar", reqId).Debug("некорректные данные Patronymic ", car.Owner.Patronymic)
				return errors.New("некорректные данные Patronymic")
			}
		}
		if car.Year != 0 {
			if car.Year < 1900 || car.Year > 2100 {
				c.logger.WithField("carCatalog.checkcar", reqId).Debug("некорректные данные Year ", car.Year)
				return errors.New("некорректные данные Year")
			}
		}
		return nil
	}

}
