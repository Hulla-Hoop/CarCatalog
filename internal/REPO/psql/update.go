package psql

import (
	"carcatalog/internal/model"
	"fmt"
	"time"
)

func (p *psql) Update(reqId string, car model.Car) (*model.CarDB, error) {
	p.logger.WithField("psql.Update", reqId).Debug("Полученные данные", car)

	carDB := car.CarToCarDB()

	query := p.getQuery(reqId, carDB)

	err := p.dB.QueryRow(query).Scan(&carDB.Id, &carDB.RegNum, &carDB.Mark, &carDB.Model, &carDB.Year, &carDB.Name, &carDB.Surname, &carDB.Patronymic, &carDB.Removed, &carDB.Created_at, &carDB.Updated_at)
	if err != nil {
		p.logger.WithField("psql.Update", reqId).Error(err)
		return nil, err
	}

	return carDB, nil
}

func (p *psql) getQuery(reqId string, carDB *model.CarDB) string {

	query := `
	UPDATE cars 
	SET `

	if carDB.RegNum != "" {
		query += fmt.Sprintf("regNum = '%s',", carDB.RegNum)
	}
	if carDB.Mark != "" {
		query += fmt.Sprintf("mark = '%s',", carDB.Mark)
	}
	if carDB.Model != "" {
		query += fmt.Sprintf("model = '%s',", carDB.Model)
	}
	if carDB.Year != 0 {
		query += fmt.Sprintf("year = %d,", carDB.Year)
	}
	if carDB.Name != "" {
		query += fmt.Sprintf("name = '%s',", carDB.Name)
	}
	if carDB.Surname != "" {
		query += fmt.Sprintf("surname = '%s',", carDB.Surname)
	}
	if carDB.Patronymic != "" {
		query += fmt.Sprintf("patronymic = '%s',", carDB.Patronymic)
	}

	query = query + fmt.Sprintf("updated_at = '%s' WHERE id = %d returning *;", carDB.Updated_at.Format(time.DateTime), carDB.Id)

	p.logger.WithField("psql.Update", reqId).Debug("Запрос на обновление", query)

	return query
}
