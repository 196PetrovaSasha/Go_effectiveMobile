package service

import (
	"context"
	"effectiveMobile/internal/fio/model"
	"effectiveMobile/internal/fio/service/create"
	"github.com/asaskevich/govalidator"
)

type CreateParams struct {
	Name       string `valid:"required"`
	Surname    string `valid:"required"`
	Patronymic string `valid:"required"`
}

func (s Service) Create(ctx context.Context, params CreateParams) (int, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return 0, error(err)
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	gender, err := create.GetGender(params.Name)
	if err != nil {
		return 0, error(err)
	}

	age, err := create.GetAge(params.Name)
	if err != nil {
		return 0, error(err)
	}

	nationality, err := create.GetNationality(params.Name)
	if err != nil {
		return 0, error(err)
	}

	entity := model.FioInformation{
		Name:        params.Name,
		Surname:     params.Surname,
		Patronymic:  params.Patronymic,
		Nationality: nationality,
		Age:         age,
		Gender:      gender,
	}
	err = s.repo.Create(ctx, &entity)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	return entity.ID, err
}
