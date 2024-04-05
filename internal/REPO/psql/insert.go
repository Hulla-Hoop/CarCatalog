package psql

import (
	"carcatalog/internal/model"
	"fmt"
	"time"
)

func (p *psql) Insert(reqId string, car model.Car) (*model.CarDB, error) {

	p.logger.WithField("psql.Insert", reqId).Debug("Полученные данные", car)

	carDB := car.CarToCarDB()

	query := fmt.Sprintf(`INSERT INTO cars (regNum, mark, model, year, name, surname, patronymic,created_at, updated_at) 
	VALUES ('%s', '%s', '%s', %d, '%s','%s','%s','%s', '%s') 
	returning *;`,
		carDB.RegNum, carDB.Mark, carDB.Model, carDB.Year, carDB.Name, carDB.Surname, carDB.Patronymic, carDB.Created_at.Format(time.DateTime), carDB.Updated_at.Format(time.DateTime))

	p.logger.WithField("psql.Insert", reqId).Debug("Запрос на вставку", query)

	err := p.dB.QueryRow(query).Scan(&carDB.Id, &carDB.RegNum, &carDB.Mark, &carDB.Model, &carDB.Year, &carDB.Name, &carDB.Surname, &carDB.Patronymic, &carDB.Removed, &carDB.Created_at, &carDB.Updated_at)
	if err != nil {
		p.logger.WithField("psql.Insert", reqId).Error(err)
		return nil, err
	}

	return carDB, nil
}
