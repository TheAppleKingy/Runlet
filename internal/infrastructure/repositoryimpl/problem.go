package repositoryimpl

import (
	"Runlet/internal/domain/entities"
	"Runlet/internal/domain/repository"
	textdata "Runlet/internal/infrastructure/text_data"
	"context"

	"github.com/doug-martin/goqu/v9"
)

type ProblemRepository struct {
	repository.ProblemRepositoryInterface
	db *goqu.Database
}

func NewProblemRepository(db *goqu.Database) *ProblemRepository {
	return &ProblemRepository{
		db: db,
	}
}

func (r ProblemRepository) GetProblem(ctx context.Context, id int) (entities.Problem, error) {
	var problem entities.Problem
	if found, err := r.db.From(textdata.ProblemTable).Select().Where(goqu.Ex{"id": id}).ScanStructContext(ctx, &problem); err != nil || !found {
		return entities.Problem{}, err
	}
	return problem, nil
}

func (r ProblemRepository) GetProblemTestCases(ctx context.Context, problemId int) (entities.TestCases, error) {
	var testCases entities.TestCases
	if found, err := r.db.From(goqu.T(textdata.ProblemTable).As("p")).
		Select(goqu.I("p.test_cases")).
		Where(goqu.I("p.id").Eq(problemId)).
		ScanValContext(ctx, &testCases); err != nil || !found {
		return entities.TestCases{}, err
	}
	return testCases, nil
}

func (r ProblemRepository) GetCourseProblems(ctx context.Context, courseId int) ([]entities.Problem, error) {
	var problems []entities.Problem
	if err := r.db.From(goqu.T(textdata.ProblemTable).As("p")).
		Select(
			goqu.I("p.id"),
			goqu.I("p.title"),
			goqu.I("p.description")).
		Where(goqu.Ex{"course_id": courseId}).ScanStructsContext(ctx, &problems); err != nil {
		return []entities.Problem{}, err
	}
	return problems, nil
}

func (r ProblemRepository) CreateProblem(ctx context.Context, title string, description string, courseId int) (entities.Problem, error) {
	return entities.Problem{}, nil
}

func (r ProblemRepository) UpdateCourseProblem(ctx context.Context, problemId int, title string, description string) error {
	return nil
}

func (r ProblemRepository) DeleteCourseProblem(ctx context.Context, problemId int) error {
	return nil
}
