package repository

import (
	"Runlet/internal/infrastructure/ent"
	"context"
)

type StudentRepositoryInterface interface {
	GetStudent(ctx context.Context, id int) (*ent.Student, error)
	GetStudentByEmail(ctx context.Context, email string) (*ent.Student, error)
	CreateStudent(ctx context.Context, name string, email string, hashedPassword string, classID int) (*ent.Student, error)
	DeleteStudent(ctx context.Context, id int) error
}
