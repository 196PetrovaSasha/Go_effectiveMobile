package repository

import (
	"context"
	"effectiveMobile/internal/fio/model"
	"effectiveMobile/pkg/db"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{Db: db}
}

func (r Repository) FindByName(ctx context.Context, name string, surname string, patronymic string) (model.FioInformation, error) {
	entity := model.FioInformation{}
	query := fmt.Sprintf(
		"SELECT * FROM fio WHERE name = $1 AND surname = $2 AND patronymic = $3 AND WHERE deleted_on IS NULL",
	)
	err := r.Db.GetContext(ctx, &entity, query, name, surname, patronymic)
	return entity, db.HandleError(err)
}

func (r Repository) Find(ctx context.Context, id int) (model.FioInformation, error) {
	entity := model.FioInformation{}
	query := fmt.Sprintf(
		"SELECT * FROM fio WHERE id = $1 AND deleted_on IS NULL",
	)
	err := r.Db.GetContext(ctx, &entity, query, id)
	return entity, db.HandleError(err)
}

func (r Repository) Create(ctx context.Context, entity *model.FioInformation) error {
	query := `INSERT INTO fio (name, surname, patronymic, nationality, age, gender)
                VALUES (:name, :surname, :patronymic, :nationality, :age, :gender) RETURNING id;`
	rows, err := r.Db.NamedQueryContext(ctx, query, entity)
	if err != nil {
		return db.HandleError(err)
	}

	for rows.Next() {
		err = rows.StructScan(entity)
		if err != nil {
			return db.HandleError(err)
		}
	}
	return db.HandleError(err)
}

func (r Repository) Update(ctx context.Context, entity model.FioInformation) error {
	query := `UPDATE fio
                SET name = :name, 
                    surname = :surname, 
                    patronymic = :patronymic, 
                    nationality = :nationality, 
                    age = :age, 
                    gender = :gender,
                	deleted_on = :deleted_on
                WHERE id = :id;`
	_, err := r.Db.NamedExecContext(ctx, query, entity)
	return db.HandleError(err)
}

func (r Repository) FindAll(ctx context.Context) ([]model.FioInformation, error) {
	var entities []model.FioInformation
	query := fmt.Sprintf(
		"SELECT * FROM fio WHERE deleted_on IS NULL",
	)
	err := r.Db.SelectContext(ctx, &entities, query)
	return entities, db.HandleError(err)
}

func (r Repository) Delete(ctx context.Context, id int) (model.FioInformation, error) {
	entity := model.FioInformation{}
	query := fmt.Sprintf(
		"DELETE FROM fio WHERE id = $1 AND deleted_on IS NULL",
	)
	err := r.Db.GetContext(ctx, &entity, query, id)
	return entity, db.HandleError(err)
}
