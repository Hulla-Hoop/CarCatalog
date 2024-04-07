package psql

import (
	"carcatalog/internal/model"
	"fmt"
)

func (p *psql) Filter(reqId string, filter map[string]string) ([]model.CarDB, error) {
	query := p.queryFilter(filter)

	p.logger.WithField("psql.Filter", reqId).Debug("Запрос на фильтрацию", query)

	rows, err := p.dB.Query(query)
	if err != nil {
		p.logger.WithField("psql.Filter", reqId).Error(err)
		return nil, err
	}
	defer rows.Close()

	var carDBSL []model.CarDB

	for rows.Next() {
		var carDB model.CarDB
		err := rows.Scan(&carDB.Id, &carDB.RegNum, &carDB.Mark, &carDB.Model, &carDB.Year, &carDB.Name, &carDB.Surname, &carDB.Patronymic, &carDB.Removed, &carDB.Created_at, &carDB.Updated_at)
		if err != nil {
			p.logger.WithField("psql.Filter", reqId).Error(err)
			return nil, err
		}
		carDBSL = append(carDBSL, carDB)
	}

	return carDBSL, nil
}

// генирирует строку запроса на основе полученной мапы которая содержит параметры фильтрации
func (p *psql) queryFilter(filter map[string]string) string {

	p.logger.WithField("psql.queryFilter", "").Debug("Полученные данные", filter)

	var operators = map[string]string{
		"eq": "=",
		"ne": "!=",
		"gt": ">",
		"ge": ">=",
		"lt": "<",
		"le": "<=",
	}

	query := `SELECT *
	FROM cars `

	if len(filter) == 0 {

		query += "WHERE removed = false"
		return query

	} else {

		query += "WHERE removed = false "

		offset, ok := filter["offset"]
		limit, es := filter["limit"]

		if ok && es {
			query += fmt.Sprintf("AND id > %s", offset)

			field, ok := filter["field"]
			value, es := filter["value"]
			operator, re := filter["operator"]
			if ok && es && re {
				if field != "id" {
					query += fmt.Sprintf(" AND %s %s '%s'", field, operators[operator], value)
				}
			}

			query += fmt.Sprintf("LIMIT %s", limit)

			return query
		} else {

			field, ok := filter["field"]
			value, es := filter["value"]
			operator, re := filter["operator"]
			if ok && es && re {
				if field != "id" {
					query += fmt.Sprintf(" AND %s %s '%s'", field, operators[operator], value)
				}
			}

			return query
		}

	}

}
