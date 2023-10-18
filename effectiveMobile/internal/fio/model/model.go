package model

import "time"

type FioInformation struct {
	ID          int        `db:"id"`
	Name        string     `db:"name"`
	Surname     string     `db:"surname"`
	Patronymic  string     `db:"patronymic"`
	Nationality string     `db:"nationality"`
	Age         int        `db:"age"`
	Gender      string     `db:"gender"`
	DeletedOn   *time.Time `db:"deleted_on"`
}
