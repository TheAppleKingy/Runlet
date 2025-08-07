package repositoryimpl

import (
	"Runlet/internal/infrastructure/ent"
)

func NewCourseRepository(dbClient *ent.Client) *CourseRepository {
	return &CourseRepository{
		client: dbClient,
	}
}
