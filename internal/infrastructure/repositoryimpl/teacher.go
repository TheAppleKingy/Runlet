package repositoryimpl

import (
	"Runlet/internal/domain/entities"
	"Runlet/internal/domain/repository"
	textdata "Runlet/internal/infrastructure/text_data"
	"context"

	"github.com/doug-martin/goqu/v9"
)

type TeacherRepository struct {
	repository.TeacherRepositoryInterface
	db *goqu.Database
}

func NewTeacherRepository(db *goqu.Database) *TeacherRepository {
	return &TeacherRepository{
		db: db,
	}
}

func (r TeacherRepository) GetTeacher(ctx context.Context, id int) (entities.Teacher, error) {
	return entities.Teacher{}, nil
}

func (r TeacherRepository) GetTeacherByEmail(ctx context.Context, email string) (entities.Teacher, error) {
	var teacher entities.Teacher
	found, err := r.db.From(textdata.TeacherTable).Select().Where(goqu.Ex{"email": email}).ScanStruct(&teacher)
	if err != nil || !found {
		return entities.Teacher{}, err
	}
	return teacher, nil
}

func (r TeacherRepository) CreateTeacher(ctx context.Context, name string, email string, hashedPassword string) (entities.Teacher, error) {
	var teacher entities.Teacher
	created, err := r.db.Insert(textdata.TeacherTable).Rows(goqu.Record{
		"name":     name,
		"email":    email,
		"password": hashedPassword,
		"is_admin": false,
	}).Returning().Executor().ScanStruct(&teacher)
	if err != nil || !created {
		return entities.Teacher{}, err
	}
	return teacher, nil
}

func (r TeacherRepository) DeleteTeacher(ctx context.Context, id int) error {
	return nil
}
