package repositoryimpl

import (
	"Runlet/internal/domain/entities"
	"Runlet/internal/domain/repository"
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
	return entities.Problem{}, nil
}

func (r ProblemRepository) GetCourseProblems(ctx context.Context, courseId int) ([]entities.Problem, error) {
	return make([]entities.Problem, 0), nil
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
