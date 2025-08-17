package tests

import (
	"Runlet/internal/application/service"
	"Runlet/internal/infrastructure/repositoryimpl"
	"Runlet/internal/infrastructure/security"
	"Runlet/internal/infrastructure/tables"
	"Runlet/internal/interfaces/http/handlers"
	"log/slog"
	"os"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exec"
	"github.com/gin-gonic/gin"
)

func setUpDb(db *goqu.Database) {
	hsh, _ := security.HashPassword("test_password")
	executors := []exec.QueryExecutor{
		db.Insert(tables.ClassTable).Rows(goqu.Record{
			"number": "111111",
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
			"amount":     0,
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

func getTestServer(db *goqu.Database) *gin.Engine {
	studentRepo := repositoryimpl.NewStudentRepository(db)
	classRepo := repositoryimpl.NewClassRepository(db)
	courseRepo := repositoryimpl.NewCourseRepository(db)
	// teacherRepo := repositoryimpl.NewTeacherRepository(db)
	// problemRepo := repositoryimpl.NewProblemRepository(db)

	studentAuthService := service.NewStudentAuthService(studentRepo, classRepo)
	studentService := service.NewStudentService(courseRepo)

	r := gin.New()
	gin.SetMode(gin.TestMode)
	testGroup := r.Group("/test")
	stGroup := testGroup.Group("/student")
	handlers.ConnectStudentHandler(stGroup, studentService, studentAuthService)
	return r
}
