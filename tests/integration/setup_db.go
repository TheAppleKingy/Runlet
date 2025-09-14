package integration

import (
	"Runlet/internal/domain/entities"
	"Runlet/internal/infrastructure/config"
	"Runlet/internal/infrastructure/security"
	"log/slog"
	"os"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exec"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var DB *goqu.Database

func setUpDb(db *goqu.Database) {
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

		db.Insert(config.Tables.Course).Rows(
			goqu.Record{
				"title":       "test_course",
				"description": "test_description",
			},
			goqu.Record{
				"title":       "test_course2",
				"description": "test_descr2",
			},
		).Executor(),

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

		db.Insert(config.Tables.Problem).Rows(
			goqu.Record{
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
			},
			goqu.Record{
				"title":       "test_problem2",
				"description": "test_pr_descr2",
				"course_id":   2,
				"test_cases": entities.TestCases{
					entities.TestCase{
						TestNum: 1,
						Input:   "3",
						Output:  "3",
					},
				},
			},
			goqu.Record{
				"title":       "test_problem3",
				"description": "test_pr_descr3",
				"course_id":   1,
				"test_cases": entities.TestCases{
					entities.TestCase{
						TestNum: 1,
						Input:   "4",
						Output:  "4",
					},
				},
			},
		).Executor(),

		db.Insert(config.Tables.ClassesCourses).Rows(goqu.Record{
			"class_id":  1,
			"course_id": 1,
		}).Executor(),

		db.Insert(config.Tables.TeachersClasses).Rows(goqu.Record{
			"teacher_id": 1,
			"class_id":   1,
		}).Executor(),

		db.Insert(config.Tables.Attempt).Rows(goqu.Record{
			"student_id": 1,
			"problem_id": 1,
			"amount":     1,
			"done":       true,
			"test_cases": []byte("[]"),
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
