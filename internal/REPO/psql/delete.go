package psql

import (
	"carcatalog/internal/model"
	"fmt"
	"time"
)

func (p *psql) Delete(reqId string, id int) (*model.CarDB, error) {

	var car model.CarDB

	time := time.Now().Format(time.DateTime)

	p.logger.WithField("psql.Delete", reqId).Debug("Полученные данные", id)

	query := fmt.Sprintf(`UPDATE cars SET removed = true, updated_at = '%s' WHERE id = %d returning *;`, time, id)

	p.logger.WithField("psql.Delete", reqId).Debug("Запрос на удаление", query)

	err := p.dB.QueryRow(query).Scan(&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Name, &car.Surname, &car.Patronymic, &car.Removed, &car.Created_at, &car.Updated_at)

	if err != nil {
		p.logger.WithField("psql.Delete", reqId).Error(err)
		return nil, err
	}
	return &car, nil
}
