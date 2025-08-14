package repository

import (
	"Runlet/internal/domain/entities"
	"context"
)

type TeacherRepositoryInterface interface {
	GetTeacher(ctx context.Context, id int) (entities.Teacher, error)
	GetTeacherByEmail(ctx context.Context, email string) (entities.Teacher, error)
	CreateTeacher(ctx context.Context, name string, email string, hashedPassword string) (entities.Teacher, error)
	DeleteTeacher(ctx context.Context, id int) error
}
