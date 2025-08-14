package fixtures

import (
	"Runlet/internal/infrastructure/security"
	"Runlet/internal/infrastructure/tables"
	"log/slog"
	"os"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exec"
)

func setUpDb(db *goqu.Database) {
	hsh, _ := security.HashPassword("test_password")
	executors := []exec.QueryExecutor{
		db.Insert(tables.ClassTable).Rows(goqu.Record{
			"num": "111111",
		}).Executor(),

		db.Insert(tables.StudentTable).Rows(goqu.Record{
			"name":     "test_student",
			"email":    "test@mail",
			"password": hsh,
			"class_id": 1,
		}).Executor(),

		db.Insert(tables.CourseTable).Rows(goqu.Record{
			"title":       "test_course",
			"description": "test_description",
		}).Executor(),

		db.Insert(tables.TeacherTable).Rows(
			goqu.Record{
				"name":     "test_teacher",
				"email":    "test_t_email",
				"password": hsh,
			},
			goqu.Record{
				"name":     "admin",
				"email":    "admin@mail",
				"password": hsh,
				"is_admin": true,
			},
		).Executor(),

		db.Insert(tables.ProblemTable).Rows(goqu.Record{
			"title":       "test_problem",
			"description": "test_pr_descr",
			"course_id":   1,
		}).Executor(),

		db.Insert("classes_courses").Rows(goqu.Record{
			"class_id":  1,
			"course_id": 1,
		}).Executor(),

		db.Insert("teachers_classes").Rows(goqu.Record{
			"teacher_id": 1,
			"class_id":   1,
		}).Executor(),

		db.Insert("attempts").Rows(goqu.Record{
			"student_id": 1,
			"problem_id": 1,
		}).Executor(),
	}
	for _, executor := range executors {
		_, err := executor.Exec()
		if err != nil {
			slog.Error("cannot setub db", "error", err)
			os.Exit(1)
		}
	}
}
