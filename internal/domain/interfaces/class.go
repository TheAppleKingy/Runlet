package interfaces

import (
	"Runlet/internal/domain/entities"
	"context"
)

type ClassRepository interface {
	GetClass(ctx context.Context, num string) (entities.Class, error)
	GetAllClasses(ctx context.Context) ([]entities.Class, error)
	CreateClass(ctx context.Context, num string) (entities.Class, error)
	DeleteClass(ctx context.Context, id int) error
}
