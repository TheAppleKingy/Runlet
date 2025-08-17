package repositoryimpl

import (
	"Runlet/internal/domain/entities"
	"Runlet/internal/domain/repository"
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
	return entities.Teacher{}, nil
}

func (r TeacherRepository) CreateTeacher(ctx context.Context, name string, email string, hashedPassword string) (entities.Teacher, error) {
	return entities.Teacher{}, nil
}

func (r TeacherRepository) DeleteTeacher(ctx context.Context, id int) error {
	return nil
}
