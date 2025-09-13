package service

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/application/service"
	"Runlet/internal/domain/entities"
	"Runlet/tests/unit/mocks"
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

// getLog replaces the default slog logger with one that writes to a buffer,
// executes the provided function f (which need to test), and returns the captured log output as a string.
// The original logger is not restored, so this helper should only be used in tests.
func getLog(t *testing.T, f func()) string {
	var buf bytes.Buffer

	logger := slog.New(slog.NewTextHandler(&buf, nil))
	slog.SetDefault(logger)

	f()

	return buf.String()
}

func TestGetCourses(t *testing.T) {
	courseRepo := mocks.NewCourseRepository(t)
	ctx := context.Background()
	student := 1
	returning := []entities.Course{{ID: 1}}

	s := service.StudentService{CourseRepository: courseRepo}
	courseRepo.On("GetAllStudentCourses", ctx, student).Return(returning, nil)

	res, err := s.GetStudentCourses(ctx, student)
	assert.NoError(t, err)
	assert.Equal(t, res, returning)
}

func TestGetProblemsOK(t *testing.T) {
	problemRepo := mocks.NewProblemRepository(t)
	courseRepo := mocks.NewCourseRepository(t)
	ctx := context.Background()
	studentId := 1
	courseId := 1
	expectedProblems := []entities.Problem{{ID: 1}}

	courseRepo.On("CheckStudent", ctx, studentId, courseId).Return(true)
	problemRepo.On("GetCourseProblems", ctx, courseId).Return(expectedProblems, nil)

	s := service.StudentService{
		CourseRepository:  courseRepo,
		ProblemRepository: problemRepo,
	}
	res, err := s.GetStudentProblems(ctx, studentId, courseId)

	assert.NoError(t, err)
	assert.Equal(t, res, expectedProblems)
}

func TestGetProblemsStudentDoesNotBelongToCourse(t *testing.T) {
	problemRepo := mocks.NewProblemRepository(t)
	courseRepo := mocks.NewCourseRepository(t)
	ctx := context.Background()
	studentId := 1
	courseId := 1
	expectedProblems := []entities.Problem{}

	courseRepo.On("CheckStudent", ctx, studentId, courseId).Return(false)

	s := service.StudentService{
		CourseRepository:  courseRepo,
		ProblemRepository: problemRepo,
	}
	res, err := s.GetStudentProblems(ctx, studentId, courseId)

	assert.Equal(t, err, errors.New("provided student does not belong to provided course"))
	assert.Equal(t, res, expectedProblems)
}

func TestSendCodeSolutionOK(t *testing.T) {
	solution := dto.CodeSolution{Lang: "foo", Code: "foo"}
	ctx := context.Background()
	studentId := 1
	problem := entities.Problem{
		ID: 1,
	}

	currentRes := entities.TestCases{
		entities.TestCase{
			TestNum: 1,
			Input:   "foo",
			Output:  "ok",
		},
	}

	attemptRepo := mocks.NewAttemptRepository(t)
	attemptRepo.On("CheckProblemIsDone", mocks.GetMockArgs(3)...).Return(false)
	attemptRepo.On("GetCurrentResults", mocks.GetMockArgs(3)...).Return(currentRes, nil)
	attemptRepo.On("AddAttempt", mocks.GetMockArgs(5)...).Return(nil)

	problemRepo := mocks.NewProblemRepository(t)
	problemRepo.On("GetProblemTestCases", mocks.GetMockArgs(2)...).Return(currentRes, nil)

	runner := mocks.NewRunner(t)
	runner.On("Run", mocks.GetMockArgs(5)...).Return(currentRes, nil)

	s := service.StudentService{
		Runner:            runner,
		AttemptRepository: attemptRepo,
		ProblemRepository: problemRepo,
	}

	log := getLog(t, func() {
		s.SendCodeSolution(ctx, studentId, problem.ID, solution)
	})

	attemptRepo.AssertExpectations(t)
	attemptRepo.AssertCalled(t, "AddAttempt", ctx, studentId, problem.ID, true, currentRes)
	problemRepo.AssertExpectations(t)
	runner.AssertExpectations(t)

	assert.Equal(t, log, "")
}

func TestSendCodeSolutionTestsMismatch(t *testing.T) {
	solution := dto.CodeSolution{Lang: "foo", Code: "foo"}
	ctx := context.Background()
	studentId := 1
	problem := entities.Problem{
		ID: 1,
	}

	currentRes := entities.TestCases{
		entities.TestCase{
			TestNum: 1,
			Input:   "foo",
			Output:  "ok",
		},
	}
	gotRes := entities.TestCases{
		{
			TestNum: 1,
			Input:   "foo",
			Output:  "not ok",
		},
	}

	attemptRepo := mocks.NewAttemptRepository(t)
	attemptRepo.On("CheckProblemIsDone", mocks.GetMockArgs(3)...).Return(false)
	attemptRepo.On("GetCurrentResults", mocks.GetMockArgs(3)...).Return(currentRes, nil)
	attemptRepo.On("AddAttempt", mocks.GetMockArgs(5)...).Return(nil)

	problemRepo := mocks.NewProblemRepository(t)
	problemRepo.On("GetProblemTestCases", mocks.GetMockArgs(2)...).Return(currentRes, nil)

	runner := mocks.NewRunner(t)
	runner.On("Run", mocks.GetMockArgs(5)...).Return(gotRes, nil)

	s := service.StudentService{
		Runner:            runner,
		AttemptRepository: attemptRepo,
		ProblemRepository: problemRepo,
	}

	log := getLog(t, func() {
		s.SendCodeSolution(ctx, studentId, problem.ID, solution)
	})

	attemptRepo.AssertExpectations(t)
	attemptRepo.AssertCalled(t, "AddAttempt", ctx, studentId, problem.ID, false, gotRes)
	problemRepo.AssertExpectations(t)
	runner.AssertExpectations(t)

	assert.Equal(t, log, "")

}

func TestSendCodeSolutionAlreadyDone(t *testing.T) {
	solution := dto.CodeSolution{Lang: "foo", Code: "foo"}
	ctx := context.Background()
	studentId := 1
	problem := entities.Problem{
		ID: 1,
	}

	attemptRepo := mocks.NewAttemptRepository(t)
	attemptRepo.On("CheckProblemIsDone", mocks.GetMockArgs(3)...).Return(true)

	problemRepo := mocks.NewProblemRepository(t)

	runner := mocks.NewRunner(t)

	s := service.StudentService{
		Runner:            runner,
		AttemptRepository: attemptRepo,
		ProblemRepository: problemRepo,
	}

	log := getLog(t, func() {
		s.SendCodeSolution(ctx, studentId, problem.ID, solution)
	})

	attemptRepo.AssertCalled(t, "CheckProblemIsDone", mocks.GetMockArgs(3)...)
	attemptRepo.AssertNumberOfCalls(t, "AddAttempt", 0)
	attemptRepo.AssertNumberOfCalls(t, "GetCurrentResults", 0)

	problemRepo.AssertNumberOfCalls(t, "GetProblemTestCases", 0)

	runner.AssertNumberOfCalls(t, "Run", 0)

	assert.Contains(t, log, fmt.Sprintf(
		"level=ERROR msg=\"problem is already done\" problem_id=%d student_id=%d runner=%s",
		problem.ID,
		studentId,
		solution.Lang,
	),
	)
}

func TestSendCodeSolutionErrorGettingCurrentResults(t *testing.T) {
	solution := dto.CodeSolution{Lang: "foo", Code: "foo"}
	ctx := context.Background()
	studentId := 1
	problem := entities.Problem{
		ID: 1,
	}

	attemptRepo := mocks.NewAttemptRepository(t)
	attemptRepo.On("CheckProblemIsDone", mocks.GetMockArgs(3)...).Return(false)
	attemptRepo.On("GetCurrentResults", mocks.GetMockArgs(3)...).Return(
		entities.TestCases{},
		errors.New("unable to get current results"),
	)

	problemRepo := mocks.NewProblemRepository(t)

	runner := mocks.NewRunner(t)

	s := service.StudentService{
		Runner:            runner,
		AttemptRepository: attemptRepo,
		ProblemRepository: problemRepo,
	}

	log := getLog(t, func() {
		s.SendCodeSolution(ctx, studentId, problem.ID, solution)
	})

	attemptRepo.AssertCalled(t, "CheckProblemIsDone", ctx, problem.ID, studentId)
	attemptRepo.AssertCalled(t, "GetCurrentResults", ctx, problem.ID, studentId)
	attemptRepo.AssertNumberOfCalls(t, "AddAttempt", 0)

	problemRepo.AssertNumberOfCalls(t, "GetProblemTestCases", 0)

	runner.AssertNumberOfCalls(t, "Run", 0)

	assert.Contains(t, log, fmt.Sprintf(
		"level=ERROR msg=\"cannot get current results\" error=\"unable to get current results\" problem_id=%d student_id=%d runner=%s\n",
		problem.ID,
		studentId,
		solution.Lang,
	),
	)
}

func TestSendCodeSolutionErrorGettingProblemTestCases(t *testing.T) {
	solution := dto.CodeSolution{Lang: "foo", Code: "foo"}
	ctx := context.Background()
	studentId := 1
	problem := entities.Problem{
		ID: 1,
	}

	attemptRepo := mocks.NewAttemptRepository(t)
	attemptRepo.On("CheckProblemIsDone", mocks.GetMockArgs(3)...).Return(false)
	attemptRepo.On("GetCurrentResults", mocks.GetMockArgs(3)...).Return(
		entities.TestCases{},
		nil,
	)

	problemRepo := mocks.NewProblemRepository(t)
	problemRepo.On("GetProblemTestCases", mocks.GetMockArgs(2)...).Return(
		entities.TestCases{},
		errors.New("unable to get problem test cases"),
	)

	runner := mocks.NewRunner(t)

	s := service.StudentService{
		Runner:            runner,
		AttemptRepository: attemptRepo,
		ProblemRepository: problemRepo,
	}

	log := getLog(t, func() {
		s.SendCodeSolution(ctx, studentId, problem.ID, solution)
	})

	attemptRepo.AssertCalled(t, "CheckProblemIsDone", ctx, problem.ID, studentId)
	attemptRepo.AssertCalled(t, "GetCurrentResults", ctx, problem.ID, studentId)
	attemptRepo.AssertNumberOfCalls(t, "AddAttempt", 0)

	problemRepo.AssertCalled(t, "GetProblemTestCases", ctx, problem.ID)

	runner.AssertNumberOfCalls(t, "Run", 0)

	assert.Contains(t, log, fmt.Sprintf(
		"level=ERROR msg=\"cannot get test cases\" error=\"unable to get problem test cases\" problem_id=%d student_id=%d runner=%s\n",
		problem.ID,
		studentId,
		solution.Lang,
	),
	)
}

func TestSendCodeSolutionNoProblemTestCases(t *testing.T) {
	solution := dto.CodeSolution{Lang: "foo", Code: "foo"}
	ctx := context.Background()
	studentId := 1
	problem := entities.Problem{
		ID: 1,
	}

	attemptRepo := mocks.NewAttemptRepository(t)
	attemptRepo.On("CheckProblemIsDone", mocks.GetMockArgs(3)...).Return(false)
	attemptRepo.On("GetCurrentResults", mocks.GetMockArgs(3)...).Return(
		entities.TestCases{},
		nil,
	)

	problemRepo := mocks.NewProblemRepository(t)
	problemRepo.On("GetProblemTestCases", mocks.GetMockArgs(2)...).Return(
		entities.TestCases{},
		nil,
	)

	runner := mocks.NewRunner(t)

	s := service.StudentService{
		Runner:            runner,
		AttemptRepository: attemptRepo,
		ProblemRepository: problemRepo,
	}

	log := getLog(t, func() {
		s.SendCodeSolution(ctx, studentId, problem.ID, solution)
	})

	attemptRepo.AssertCalled(t, "CheckProblemIsDone", ctx, problem.ID, studentId)
	attemptRepo.AssertCalled(t, "GetCurrentResults", ctx, problem.ID, studentId)
	attemptRepo.AssertNumberOfCalls(t, "AddAttempt", 0)

	problemRepo.AssertCalled(t, "GetProblemTestCases", ctx, problem.ID)

	runner.AssertNumberOfCalls(t, "Run", 0)

	assert.Contains(t, log, fmt.Sprintf(
		"level=ERROR msg=\"cannot get test cases\" error=<nil> problem_id=%d student_id=%d runner=%s\n",
		problem.ID,
		studentId,
		solution.Lang,
	),
	)
}

func TestSendCodeSolutionErrorCallGRPC(t *testing.T) {
	solution := dto.CodeSolution{Lang: "foo", Code: "foo"}
	ctx := context.Background()
	studentId := 1
	problem := entities.Problem{
		ID: 1,
	}
	currentRes := entities.TestCases{entities.TestCase{
		TestNum: 1,
		Input:   "foo",
		Output:  "ok",
	},
	}

	attemptRepo := mocks.NewAttemptRepository(t)
	attemptRepo.On("CheckProblemIsDone", mocks.GetMockArgs(3)...).Return(false)
	attemptRepo.On("GetCurrentResults", mocks.GetMockArgs(3)...).Return(
		currentRes,
		nil,
	)
	attemptRepo.On("AddAttempt", mocks.GetMockArgs(5)...).Return(nil)

	problemRepo := mocks.NewProblemRepository(t)
	problemRepo.On("GetProblemTestCases", mocks.GetMockArgs(2)...).Return(
		currentRes,
		nil,
	)

	runner := mocks.NewRunner(t)
	runner.On("Run", mocks.GetMockArgs(5)...).Return(
		entities.TestCases{},
		errors.New("error calling grpc"),
	)

	s := service.StudentService{
		Runner:            runner,
		AttemptRepository: attemptRepo,
		ProblemRepository: problemRepo,
	}

	log := getLog(t, func() {
		s.SendCodeSolution(ctx, studentId, problem.ID, solution)
	})

	attemptRepo.AssertCalled(t, "CheckProblemIsDone", ctx, problem.ID, studentId)
	attemptRepo.AssertCalled(t, "GetCurrentResults", ctx, problem.ID, studentId)
	attemptRepo.AssertCalled(t, "AddAttempt", ctx, studentId, problem.ID, false, entities.TestCases{})

	problemRepo.AssertCalled(t, "GetProblemTestCases", ctx, problem.ID)

	runner.AssertCalled(t, "Run", ctx, studentId, problem.ID, solution, []dto.RunData{
		{
			TestNum: currentRes[0].TestNum,
			Input:   currentRes[0].Input,
		},
	},
	)
	assert.Contains(t, log, fmt.Sprintf(
		"level=ERROR msg=\"error remote running code\" error=\"error calling grpc\" problem_id=%d student_id=%d runner=%s\n",
		problem.ID,
		studentId,
		solution.Lang,
	))
}
