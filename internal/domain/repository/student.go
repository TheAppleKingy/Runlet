package repository

import (
	"Runlet/internal/domain/entities"
	"context"
)

type StudentRepositoryInterface interface {
	GetStudent(ctx context.Context, id int) (entities.Student, error)
	GetStudentByEmail(ctx context.Context, email string) (entities.Student, error)
	CreateStudent(ctx context.Context, name string, email string, hashedPassword string, classID int) (entities.Student, error)
	DeleteStudent(ctx context.Context, id int) error
}
