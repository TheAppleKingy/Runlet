package repositoryimpl

import (
	"Runlet/internal/domain/entities"
	"Runlet/internal/domain/repository"
	textdata "Runlet/internal/infrastructure/text_data"
	"context"

	"github.com/doug-martin/goqu/v9"
)

type ClassRepository struct {
	repository.ClassRepositoryInterface
	db *goqu.Database
}

func NewClassRepository(db *goqu.Database) *ClassRepository {
	return &ClassRepository{
		db: db,
	}
}

func (r ClassRepository) GetClass(ctx context.Context, num string) (entities.Class, error) {
	var class entities.Class
	found, err := r.db.From(textdata.ClassTable).Select().Where(goqu.Ex{"number": num}).ScanStruct(&class)
	if err != nil || !found {
		return entities.Class{}, err
	}
	return class, nil
}

func (r ClassRepository) CreateClass(ctx context.Context, num string) (entities.Class, error) {
	var class entities.Class
	created, err := r.db.Insert(textdata.ClassTable).Rows(goqu.Record{
		"number": num,
	}).Returning().Executor().ScanStruct(&class)
	if err != nil || !created {
		return entities.Class{}, err
	}
	return class, nil
}

func (r ClassRepository) DeleteClass(ctx context.Context, id int) error {
	_, err := r.db.Delete(textdata.ClassTable).Where(goqu.Ex{"id": id}).Executor().Exec()
	return err
}
