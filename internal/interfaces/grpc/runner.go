package grpc

import (
	"Runlet/internal/application/dto"
	runner "Runlet/internal/infrastructure/proto"
	textdata "Runlet/internal/infrastructure/text_data"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Runner struct {
	Lang       string
	grpcClient runner.RunnerClient
}

func NewRunner(lang string) (*Runner, error) {
	runnerUrl, ok := textdata.Runners[lang]
	if !ok {
		return nil, fmt.Errorf("runner for lang %s does not exist", lang)
	}
	conn, err := grpc.NewClient(runnerUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	cl := runner.NewRunnerClient(conn)
	return &Runner{
		Lang:       lang,
		grpcClient: cl,
	}, nil
}

func (r Runner) Run(ctx context.Context, studentId int, problemid int, code string, cases []dto.RunTestData) (*runner.RunCodeResponse, error) {
	casesBytes, err := json.Marshal(cases)
	if err != nil {
		return nil, err
	}
	req := runner.RunCodeRequest{
		Student: int32(studentId),
		Problem: int32(problemid),
		Lang:    r.Lang,
		Code:    code,
		Cases:   casesBytes,
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return r.grpcClient.RunCode(ctx, &req)
}
