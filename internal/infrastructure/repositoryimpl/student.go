package repositoryimpl

import (
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/infrastructure/ent/student"
	"context"
)

type StudentRepository struct {
	repository.StudentRepositoryInterface
	client *ent.Client
}

func (r StudentRepository) GetStudent(ctx context.Context, id int) (*ent.Student, error) {
	return r.client.Student.Get(ctx, id)
}

func (r StudentRepository) GetStudentByEmail(ctx context.Context, email string) (*ent.Student, error) {
	return r.client.Student.Query().Where(student.EmailEQ(email)).Only(ctx)
}

func (r StudentRepository) CreateStudent(ctx context.Context, name string, email string, hashedPassword string, classID int) (*ent.Student, error) {
	return r.client.Student.Create().SetName(name).SetEmail(email).SetPassword(hashedPassword).SetClassID(classID).Save(ctx)
}

func (r StudentRepository) DeleteStudent(ctx context.Context, id int) error {
	return r.client.Student.DeleteOneID(id).Exec(ctx)
}
