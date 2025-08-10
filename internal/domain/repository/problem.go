package repository

import (
	"Runlet/internal/infrastructure/ent"
	"context"
)

type ProblemRepositoryInterface interface {
	GetProblem(ctx context.Context, id int) (*ent.Problem, error)
	GetCourseProblems(ctx context.Context, courseId int) ([]*ent.Problem, error)
	CreateProblem(ctx context.Context, title string, description string, courseId int) (ent.Problem, error)
	UpdateProblem(ctx context.Context, problemId int, title string, description string) error
	DeleteProblem(ctx context.Context, problemId int) error
}
