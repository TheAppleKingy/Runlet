package implementations

import (
	"Runlet/internal/domain/entities"
	"Runlet/internal/infrastructure/config"
	"context"

	"github.com/doug-martin/goqu/v9"
)

type ClassRepository struct {
	db *goqu.Database
}

func NewClassRepository(db *goqu.Database) *ClassRepository {
	return &ClassRepository{
		db: db,
	}
}

func (r ClassRepository) GetClass(ctx context.Context, num string) (entities.Class, error) {
	var class entities.Class
	found, err := r.db.From(config.Tables.Class).Select().Where(goqu.Ex{"number": num}).ScanStruct(&class)
	if err != nil || !found {
		return entities.Class{}, err
	}
	return class, nil
}

func (r ClassRepository) CreateClass(ctx context.Context, num string) (entities.Class, error) {
	var class entities.Class
	created, err := r.db.Insert(config.Tables.Class).Rows(goqu.Record{
		"number": num,
	}).Returning().Executor().ScanStruct(&class)
	if err != nil || !created {
		return entities.Class{}, err
	}
	return class, nil
}

func (r ClassRepository) DeleteClass(ctx context.Context, id int) error {
	_, err := r.db.Delete(config.Tables.Class).Where(goqu.Ex{"id": id}).Executor().Exec()
	return err
}

func (r ClassRepository) GetAllClasses(ctx context.Context) ([]entities.Class, error) {
	return []entities.Class{}, nil
}
