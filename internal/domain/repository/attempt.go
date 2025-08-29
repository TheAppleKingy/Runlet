package repository

import (
	"Runlet/internal/domain/entities"
	"context"
)

type AttemptRepositoryInteface interface {
	AddAttepmt(ctx context.Context, studentid int, problemId int, done bool, lastTests entities.TestCases) error
	CheckProblemIsDone(ctx context.Context, problemId int, studentId int) bool
	GetCurrentResults(ctx context.Context, problemId int, studentId int) (entities.TestCases, error)
}
