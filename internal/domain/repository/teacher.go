package repository

import (
	"Runlet/internal/infrastructure/ent"
	"context"
)

type TeacherRepositoryInterface interface {
	GetTeacher(ctx context.Context, id int) (*ent.Teacher, error)
	GetTeacherByEmail(ctx context.Context, email string) (*ent.Teacher, error)
	CreateTeacher(ctx context.Context, name string, email string, hashedPassword string) (*ent.Teacher, error)
	DeleteTeacher(ctx context.Context, id int) error
}
