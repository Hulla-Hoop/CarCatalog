package psql

import (
	"carcatalog/internal/model"
	"fmt"
)

func (p *psql) Insert(reqId string, car model.Car) (*model.CarDB, error) {

	p.logger.WithField("psql.Insert", reqId).Debug("Полученные данные", car)

	var carDB model.CarDB

	query := fmt.Sprintf(`INSERT INTO cars (regNum, mark, model, year, name, surname, patronymic) 
	VALUES ('%s', '%s', '%s', %d, '%s','%s','%s') 
	returning *;`,
		car.RegNum, car.Mark, car.Model, car.Year, car.Owner.Name, car.Owner.Surname, car.Owner.Patronymic)

	err := p.dB.QueryRow(query).Scan(&carDB.Id, &carDB.RegNum, &carDB.Mark, &carDB.Model, &carDB.Year, &carDB.Name, &carDB.Surname, &carDB.Patronymic)
	if err != nil {
		p.logger.WithField("psql.Insert", reqId).Error(err)
		return nil, err
	}
	return &carDB, nil
}
