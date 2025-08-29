package service

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/domain/entities"
	"Runlet/internal/domain/repository"
	"Runlet/internal/interfaces/grpc"
	"context"
	"encoding/json"
	"log/slog"
)

type StudentService struct {
	CourseRepository  repository.CourseRepositoryInterface
	ProblemRepository repository.ProblemRepositoryInterface
	AttemptRepository repository.AttemptRepositoryInteface
}

func NewStudentService(
	courseRepo repository.CourseRepositoryInterface,
	problemRepo repository.ProblemRepositoryInterface,
	attemptRepo repository.AttemptRepositoryInteface) *StudentService {
	return &StudentService{
		CourseRepository:  courseRepo,
		ProblemRepository: problemRepo,
		AttemptRepository: attemptRepo,
	}
}

func (s StudentService) GetStudentCourses(ctx context.Context, studentId int) ([]entities.Course, error) {
	return s.CourseRepository.GetAllStudentCourses(ctx, studentId)
}

func (s StudentService) GetStudentProblems(ctx context.Context, studentId int, courseId int) ([]entities.Problem, error) {
	return s.ProblemRepository.GetCourseProblems(ctx, courseId)
}

func (s StudentService) SendCodeSolution(ctx context.Context, studentId int, problemId int, data dto.CodeSolution) {
	done := s.AttemptRepository.CheckProblemIsDone(ctx, problemId, studentId)
	if done {
		slog.Error("problem is already done", "problem_id", problemId, "student_id", studentId, "runner", data.Lang)
		return
	}

	results, err := s.AttemptRepository.GetCurrentResults(ctx, problemId, studentId)
	if err != nil {
		slog.Error("cannot get current results", "problem_id", problemId, "student_id", studentId, "runner", data.Lang)
	}
	defer func() {
		//nolint:errcheck
		s.AttemptRepository.AddAttepmt(ctx, studentId, problemId, done, results)
	}()

	cases, err := s.ProblemRepository.GetProblemTestCases(ctx, problemId)
	if err != nil || len(cases) == 0 {
		slog.Error("cannot get test cases", "problem_id", problemId, "student_id", studentId, "runner", data.Lang)
		return
	}
	testCasesMap := make(map[int]entities.TestCase)
	var testsData []dto.RunTestData
	for _, testCase := range cases {
		testCasesMap[testCase.TestNum] = testCase
		testsData = append(testsData, dto.RunTestData{TestNum: testCase.TestNum, Input: testCase.Input})
	}

	runner, err := grpc.NewRunner(data.Lang)
	if err != nil {
		slog.Error("cannot get grpc client", "error", err, "problem_id", problemId, "student_id", studentId, "runner", data.Lang)
		return
	}
	resp, err := runner.Run(ctx, studentId, problemId, data.Code, testsData)
	if err != nil {
		slog.Error("cannot call grpc method", "error", err, "problem_id", problemId, "student_id", studentId, "runner", data.Lang)
		return
	}
	if err := json.Unmarshal(resp.Results, &results); err != nil {
		slog.Error("cannot decode grpc response", "problem_id", problemId, "student_id", studentId, "runner", data.Lang)
		return
	}

	testsPassed := true
	for _, caseRes := range results {
		if caseRes.Output != testCasesMap[caseRes.TestNum].Output {
			testsPassed = false
			break
		}
	}
	if testsPassed {
		done = true
	}
}
