package service

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/domain/entities"
	"Runlet/internal/domain/interfaces"
	"errors"

	"context"
	"log/slog"
)

type StudentService struct {
	CourseRepository  interfaces.CourseRepository
	ProblemRepository interfaces.ProblemRepository
	AttemptRepository interfaces.AttemptRepository
	Runner            interfaces.Runner
}

func NewStudentService(
	courseRepo interfaces.CourseRepository,
	problemRepo interfaces.ProblemRepository,
	attemptRepo interfaces.AttemptRepository,
	runner interfaces.Runner) *StudentService {
	return &StudentService{
		CourseRepository:  courseRepo,
		ProblemRepository: problemRepo,
		AttemptRepository: attemptRepo,
		Runner:            runner,
	}
}

func (s StudentService) GetStudentCourses(ctx context.Context, studentId int) ([]entities.Course, error) {
	return s.CourseRepository.GetAllStudentCourses(ctx, studentId)
}

func (s StudentService) GetStudentProblems(ctx context.Context, studentId int, courseId int) ([]entities.Problem, error) {
	if !s.CourseRepository.CheckStudent(ctx, studentId, courseId) {
		return []entities.Problem{}, errors.New("provided student does not belong to provided course")
	}
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
		slog.Error("cannot get current results", "error", err, "problem_id", problemId, "student_id", studentId, "runner", data.Lang)
		return
	}

	cases, err := s.ProblemRepository.GetProblemTestCases(ctx, problemId)
	if err != nil || len(cases) == 0 {
		slog.Error("cannot get test cases", "error", err, "problem_id", problemId, "student_id", studentId, "runner", data.Lang)
		return
	}

	defer func() {
		//nolint:errcheck
		s.AttemptRepository.AddAttempt(ctx, studentId, problemId, done, results)
	}()

	testCasesMap := make(map[int]entities.TestCase)
	var testsData []dto.RunData
	for _, testCase := range cases {
		testCasesMap[testCase.TestNum] = testCase
		testsData = append(testsData, dto.RunData{TestNum: testCase.TestNum, Input: testCase.Input})
	}

	results, err = s.Runner.Run(ctx, studentId, problemId, data, testsData)
	if err != nil {
		slog.Error("error remote running code", "error", err, "problem_id", problemId, "student_id", studentId, "runner", data.Lang)
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
