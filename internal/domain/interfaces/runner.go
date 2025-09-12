package interfaces

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/domain/entities"
	"context"
)

// Runner is interface for a wrapper for the proto generated grpc_interfaces.RunnerClient
type Runner interface {
	Run(ctx context.Context, studentId int, problemid int, solution dto.CodeSolution, cases []dto.RunData) (entities.TestCases, error)
}
