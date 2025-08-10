package repositoryimpl

import (
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/infrastructure/ent/problem"
	"context"
)

type ProblemRepository struct {
	repository.ProblemRepositoryInterface
	client *ent.Client
}

func (r ProblemRepository) GetProblem(ctx context.Context, id int) (*ent.Problem, error) {
	return r.client.Problem.Get(ctx, id)
}

func (r ProblemRepository) GetCourseProblems(ctx context.Context, courseId int) ([]*ent.Problem, error) {
	return r.client.Problem.Query().Where(problem.CourseIDEQ(courseId)).All(ctx)
}

func (r ProblemRepository) CreateProblem(ctx context.Context, title string, description string, courseId int) (*ent.Problem, error) {
	return r.client.Problem.Create().SetTitle(title).SetDescription(description).SetCourseID(courseId).Save(ctx)
}

func (r ProblemRepository) UpdateCourseProblem(ctx context.Context, problemId int, title string, description string) error {
	query := r.client.Problem.Update().Where(problem.IDEQ(problemId))
	if title != "" {
		query = query.SetTitle(title)
	}
	if description != "" {
		query = query.SetDescription(description)
	}
	return query.Exec(ctx)
}

func (r ProblemRepository) DeleteCourseProblem(ctx context.Context, problemId int) error {
	return r.client.Problem.DeleteOneID(problemId).Exec(ctx)
}
