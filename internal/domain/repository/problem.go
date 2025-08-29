package repository

import (
	"Runlet/internal/domain/entities"
	"context"
)

type ProblemRepositoryInterface interface {
	GetProblem(ctx context.Context, id int) (entities.Problem, error)
	GetProblemTestCases(ctx context.Context, problemId int) (entities.TestCases, error)
	GetCourseProblems(ctx context.Context, courseId int) ([]entities.Problem, error)
	CreateProblem(ctx context.Context, title string, description string, courseId int) (entities.Problem, error)
	UpdateProblem(ctx context.Context, problemId int, title string, description string) error
	DeleteProblem(ctx context.Context, problemId int) error
}
