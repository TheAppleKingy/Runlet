package service

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/ent"
	"context"
)

type CourseService struct {
	courseRepo repository.CourseRepositoryInterface
}

func (cs *CourseService) GetAllCourses(ctx context.Context) ([]*ent.Course, error) {
	return cs.courseRepo.GetAllCourses(ctx)
}

func (cs *CourseService) CreateCourse(ctx context.Context, data dto.CourseCreateDTO) (*ent.Course, error) {
	return cs.courseRepo.CreateCourse(ctx, data.Title, data.Description, data.ClassesIds)
}
