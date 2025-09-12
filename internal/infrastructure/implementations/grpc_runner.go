package implementations

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/config"
	"Runlet/internal/domain/entities"
	"Runlet/internal/domain/interfaces"
	grpc_interfaces "Runlet/internal/infrastructure/proto"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// newGRPCClient creates connection with runner container according to provided lang, creates and returns gRPC client
func newGRPCClient(lang string) (grpc_interfaces.RunnerClient, error) {
	runnerUrl, ok := config.Runners[lang]
	if !ok {
		return nil, fmt.Errorf("runner for lang %s did not registered", lang)
	}
	conn, err := grpc.NewClient(runnerUrl, grpc.WithTransportCredentials(insecure.NewCredentials())) // have to set transport security!!!
	if err != nil {
		return nil, err
	}
	return grpc_interfaces.NewRunnerClient(conn), nil
}

// GRPCRunner is implementaion of a wrapper for the proto generated grpc_interfaces.RunnerClient
type GRPCRunner struct {
	interfaces.Runner
}

func NewGRPCRunner() *GRPCRunner {
	return &GRPCRunner{}
}

func (r GRPCRunner) Run(ctx context.Context, studentId int, problemid int, solution dto.CodeSolution, cases []dto.RunData) (entities.TestCases, error) {
	casesBytes, err := json.Marshal(cases)
	if err != nil {
		return entities.TestCases{}, err
	}
	req := grpc_interfaces.RunCodeRequest{
		Student: int32(studentId),
		Problem: int32(problemid),
		Lang:    solution.Lang,
		Code:    solution.Code,
		Cases:   casesBytes,
	}

	var results entities.TestCases
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	cl, err := newGRPCClient(solution.Lang)
	if err != nil {
		return entities.TestCases{}, err
	}
	resp, err := cl.RunCode(ctx, &req)
	if err != nil {
		return entities.TestCases{}, err
	}
	if err := json.Unmarshal(resp.Results, &results); err != nil {
		return entities.TestCases{}, err
	}
	return results, nil
}
