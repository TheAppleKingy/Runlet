package course

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/ent"
	"context"
)

type CourseService struct {
	CourseRepo repository.CourseRepositoryInterface
}

func (cs *CourseService) GetAllCourses(ctx context.Context) ([]*ent.Course, error) {
	return cs.CourseRepo.GetAllCourses(ctx)
}

func (cs *CourseService) CreateCourse(ctx context.Context, data dto.CourseCreateDTO) (*ent.Course, error) {
	return cs.CourseRepo.CreateCourse(ctx, data.Title, data.Description, data.ClassesIds, data.TeachersIds)
}
