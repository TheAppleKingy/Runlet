package repositoryimpl

import (
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/infrastructure/ent/class"
	"context"
)

type ClassRepository struct {
	repository.ClassRepositoryInterface
	client *ent.Client
}

func (r ClassRepository) GetClass(ctx context.Context, num string) (*ent.Class, error) {
	return r.client.Class.Query().Where(class.NumberEQ(num)).Only(ctx)
}

func (r ClassRepository) CreateClass(ctx context.Context, num string) (*ent.Class, error) {
	return r.client.Class.Create().SetNumber(num).Save(ctx)
}

func (r ClassRepository) DeleteClass(ctx context.Context, id int) error {
	return r.client.Class.DeleteOneID(id).Exec(ctx)
}
