package model

import "time"

type Car struct {
	Id int

	RegNum string `json:"regNum"`

	Mark string `json:"mark"`

	Model string `json:"model"`

	Year int32 `json:"year,omitempty"`

	Owner *People `json:"owner"`
}

type People struct {
	Name string `json:"name"`

	Surname string `json:"surname"`

	Patronymic string `json:"patronymic,omitempty"`
}

func (c Car) CarToCarDB() *CarDB {

	return &CarDB{
		RegNum:     c.RegNum,
		Mark:       c.Mark,
		Model:      c.Model,
		Year:       c.Year,
		Name:       c.Owner.Name,
		Surname:    c.Owner.Surname,
		Patronymic: c.Owner.Patronymic,
		Removed:    false,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

}

type RegNums struct {
	RegNums []string `json:"regNums"`
}
