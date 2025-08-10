package repositoryimpl

import (
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/infrastructure/ent/teacher"
	"context"
)

type TeacherRepository struct {
	repository.TeacherRepositoryInterface
	client *ent.Client
}

func (r TeacherRepository) GetTeacher(ctx context.Context, id int) (*ent.Teacher, error) {
	return r.client.Teacher.Get(ctx, id)
}

func (r TeacherRepository) GetTeacherByEmail(ctx context.Context, email string) (*ent.Teacher, error) {
	return r.client.Teacher.Query().Where(teacher.EmailEQ(email)).Only(ctx)
}

func (r TeacherRepository) CreateTeacher(ctx context.Context, name string, email string, hashedPassword string) (*ent.Teacher, error) {
	return r.client.Teacher.Create().SetName(name).SetEmail(email).SetPassword(hashedPassword).Save(ctx)
}

func (r TeacherRepository) DeleteTeacher(ctx context.Context, id int) error {
	return r.client.Teacher.DeleteOneID(id).Exec(ctx)
}
