package service

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/domain/entities"
	"Runlet/internal/domain/repository"
	"context"
	"fmt"
)

type CourseService struct {
	CourseRepo repository.CourseRepositoryInterface
}

func (cs *CourseService) GetAllCourses(ctx context.Context) ([]entities.Course, error) {
	return cs.CourseRepo.GetAllCourses(ctx)

}

func (cs *CourseService) CreateCourse(ctx context.Context, data dto.CourseCreate) (entities.Course, error) {
	created, err := cs.CourseRepo.CreateCourse(ctx, data.Title, data.Description, data.ClassesIds, data.TeachersIds)
	if err != nil {
		return entities.Course{}, fmt.Errorf("unable to create course: %v", err)
	}
	return created, nil
}
