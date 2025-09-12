package fixtures

import (
	"Runlet/internal/config"
	"Runlet/internal/domain/entities"
	"Runlet/internal/infrastructure/security"
	"log/slog"
	"os"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exec"
)

func SetUpDb(db *goqu.Database) {
	hsh, _ := security.HashPassword("test_password")
	executors := []exec.QueryExecutor{
		db.Insert(config.Tables.Class).Rows(goqu.Record{
			"number": "111111",
		}).Executor(),

		db.Insert(config.Tables.Student).Rows(goqu.Record{
			"name":     "test_student",
			"email":    "test@mail",
			"password": hsh,
			"class_id": 1,
		}).Executor(),

		db.Insert(config.Tables.Course).Rows(goqu.Record{
			"title":       "test_course",
			"description": "test_description",
		}).Executor(),

		db.Insert(config.Tables.Teacher).Rows(
			goqu.Record{
				"name":     "test_teacher",
				"email":    "test_t@mail",
				"password": hsh,
				"is_admin": false,
			},
			goqu.Record{
				"name":     "admin",
				"email":    "admin@mail",
				"password": hsh,
				"is_admin": true,
			},
		).Executor(),

		db.Insert(config.Tables.Problem).Rows(goqu.Record{
			"title":       "test_problem",
			"description": "test_pr_descr",
			"course_id":   1,
			"test_cases": entities.TestCases{
				entities.TestCase{
					TestNum: 1,
					Input:   "2",
					Output:  "2",
				},
			},
		}).Executor(),

		db.Insert("classes_courses").Rows(goqu.Record{
			"class_id":  1,
			"course_id": 1,
		}).Executor(),

		db.Insert("teachers_classes").Rows(goqu.Record{
			"teacher_id": 1,
			"class_id":   1,
		}).Executor(),
	}
	for idx, executor := range executors {
		_, err := executor.Exec()
		if err != nil {
			slog.Error("cannot setub db", "error", err, "executor num", idx+1)
			os.Exit(1)
		}
	}
}
