package repositoryimpl

import (
	"Runlet/internal/domain/entities"
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/tables"
	"context"

	"github.com/doug-martin/goqu/v9"
)

type StudentRepository struct {
	repository.StudentRepositoryInterface
	db *goqu.Database
}

func NewStudentRepository(db *goqu.Database) *StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

func (r StudentRepository) GetStudent(ctx context.Context, id int) (entities.Student, error) {
	var student entities.Student
	found, err := r.db.From(tables.StudentTable).Select().Where(goqu.Ex{"id": id}).ScanStruct(&student)
	if err != nil || !found {
		return entities.Student{}, err
	}
	return student, nil
}

func (r StudentRepository) GetStudentByEmail(ctx context.Context, email string) (entities.Student, error) {
	var student entities.Student
	found, err := r.db.From(tables.StudentTable).Select().Where(goqu.Ex{"email": email}).ScanStruct(&student)
	if err != nil || !found {
		return entities.Student{}, err
	}
	return student, nil
}

func (r StudentRepository) CreateStudent(ctx context.Context, name string, email string, hashedPassword string, classID int) (entities.Student, error) {
	var student entities.Student
	created, err := r.db.Insert(tables.StudentTable).Rows(goqu.Record{
		"name":     name,
		"email":    email,
		"password": hashedPassword,
		"class_id": classID,
	}).Returning().Executor().ScanStruct(&student)
	if err != nil || !created {
		return entities.Student{}, err
	}
	return student, nil
}

func (r StudentRepository) DeleteStudent(ctx context.Context, id int) error {
	return nil
}
