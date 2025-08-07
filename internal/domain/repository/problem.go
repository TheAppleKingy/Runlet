package repository

import "context"

type ProblemsRepositoryInterface interface {
	GetProblemById(ctx context.Context)
}
