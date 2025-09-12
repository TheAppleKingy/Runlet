package implementations

import (
	"Runlet/internal/config"
	"Runlet/internal/domain/entities"
	"Runlet/internal/domain/interfaces"
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

type AttemptRepository struct {
	interfaces.AttemptRepository
	db *goqu.Database
}

func NewAttemptRepository(db *goqu.Database) *AttemptRepository {
	return &AttemptRepository{
		db: db,
	}
}

func (r AttemptRepository) AddAttempt(ctx context.Context, studentid int, problemId int, done bool, lastTests entities.TestCases) error {
	_, err := r.db.Insert(config.Tables.Attempt).Rows(
		goqu.Record{
			"student_id": studentid,
			"problem_id": problemId,
			"amount":     1,
			"done":       done,
			"test_cases": lastTests,
		},
	).OnConflict(
		goqu.DoUpdate(
			"student_id, problem_id",
			goqu.Record{
				"amount":     goqu.L(fmt.Sprintf("%s.amount+1", config.Tables.Attempt)),
				"done":       done,
				"test_cases": lastTests,
			},
		),
	).Executor().ExecContext(ctx)
	return err
}

func (r AttemptRepository) CheckProblemIsDone(ctx context.Context, problemId int, studentId int) bool {
	var done bool
	if _, err := r.db.From(goqu.T(config.Tables.Attempt).As("a")).
		Select(goqu.I("a.done")).
		Where(
			goqu.I("a.problem_id").Eq(problemId),
			goqu.I("a.student_id").Eq(studentId)).
		ScanValContext(ctx, &done); err != nil {
		return false
	}
	return done
}

func (r AttemptRepository) GetCurrentResults(ctx context.Context, problemId int, studentId int) (entities.TestCases, error) {
	var results entities.TestCases
	if found, err := r.db.From(goqu.T(config.Tables.Attempt).As("a")).
		Select(goqu.I("a.test_cases")).
		Where(
			goqu.I("a.problem_id").Eq(problemId),
			goqu.I("a.student_id").Eq(studentId)).
		ScanValContext(ctx, &results); err != nil || !found {
		return entities.TestCases{}, err
	}
	return results, nil
}
