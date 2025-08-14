package repositoryimpl

import (
	"Runlet/internal/domain/entities"
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/tables"
	"context"
	"fmt"

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
	found, err := r.db.From(tables.ClassTable).Select().Where(goqu.Ex{"number": num}).ScanStruct(&class)
	if err != nil || !found {
		fmt.Println(err, "err here")
		return entities.Class{}, err
	}
	return class, nil
}

func (r ClassRepository) CreateClass(ctx context.Context, num string) (entities.Class, error) {
	var class entities.Class
	created, err := r.db.Insert(tables.ClassTable).Rows(goqu.Record{
		"number": num,
	}).Returning().Executor().ScanStruct(&class)
	if err != nil || !created {
		return entities.Class{}, err
	}
	return class, nil
}

func (r ClassRepository) DeleteClass(ctx context.Context, id int) error {
	_, err := r.db.Delete(tables.ClassTable).Where(goqu.Ex{"id": id}).Executor().Exec()
	return err
}
