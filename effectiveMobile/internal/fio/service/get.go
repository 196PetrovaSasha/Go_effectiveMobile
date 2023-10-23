package service

import (
	"context"
	"effectiveMobile/internal/fio/model"
	"effectiveMobile/pkg/db"
	"errors"
	"github.com/fir1/rest-api/pkg/erru"
)

func (s Service) Get(ctx context.Context, id int) (model.FioInformation, error) {
	todo, err := s.repo.Find(ctx, id)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return model.FioInformation{}, erru.ErrArgument{errors.New("todo object not found")}
	default:
		return model.FioInformation{}, err
	}
	return todo, nil
}
