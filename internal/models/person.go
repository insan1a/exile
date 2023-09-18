package models

import "time"

type Person struct {
	ID          string    `db:"id"`
	Name        string    `db:"name" validate:"required,alpha"`
	Surname     string    `db:"surname" validate:"required,alpha"`
	Patronymic  string    `db:"patronymic" validate:"omitempty,alpha"`
	Age         int       `db:"age"`
	Gender      string    `db:"gender"`
	Nationality string    `db:"nationality"`
	CreatedOn   time.Time `db:"created_on"`
}
