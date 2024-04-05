package model

import "time"

type CarDB struct {
	Id int

	RegNum string `json:"regNum"`

	Mark string `json:"mark"`

	Model string `json:"model"`

	Year int32 `json:"year,omitempty"`

	Name string `json:"name"`

	Surname string `json:"surname"`

	Patronymic string `json:"patronymic,omitempty"`

	Removed bool `json:"removed"`

	Created_at time.Time `json:"created_at"`

	Updated_at time.Time `json:"updated_at"`
}

func (c *CarDB) CarToCar() *Car {

	return &Car{
		Id:     c.Id,
		RegNum: c.RegNum,
		Mark:   c.Mark,
		Model:  c.Model,
		Year:   c.Year,
		Owner: &People{
			Name:       c.Name,
			Surname:    c.Surname,
			Patronymic: c.Patronymic,
		},
	}
}
