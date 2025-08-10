package repository

import (
	"Runlet/internal/infrastructure/ent"
	"context"
)

type ClassRepositoryInterface interface {
	GetClass(ctx context.Context, num string) (*ent.Class, error)
	CreateClass(ctx context.Context, num string) (*ent.Class, error)
	DeleteClass(ctx context.Context, id int) error
}
