package service

import (
	"context"
	"github.com/asaskevich/govalidator"
	"github.com/fir1/rest-api/pkg/erru"
)

type UpdateParams struct {
	ID          int `valid:"required"`
	Name        *string
	Surname     *string
	Patronymic  *string
	Nationality *string
	Age         *int
	Gender      *string
}

func (s Service) Update(ctx context.Context, params UpdateParams) error {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return erru.ErrArgument{Wrapped: err}
	}

	todo, err := s.Get(ctx, params.ID)
	if err != nil {
		return err
	}

	if params.Name != nil {
		todo.Name = *params.Name
	}
	if params.Patronymic != nil {
		todo.Patronymic = *params.Patronymic
	}
	if params.Surname != nil {
		todo.Surname = *params.Surname
	}

	if params.Nationality != nil {
		todo.Nationality = *params.Nationality
	}

	if params.Age != nil {
		todo.Age = *params.Age
	}

	if params.Gender != nil {
		todo.Gender = *params.Gender
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = s.repo.Update(ctx, todo)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
