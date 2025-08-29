package tests

import (
	"Runlet/internal/application/service"
	"Runlet/internal/domain/entities"
	"Runlet/internal/infrastructure/repositoryimpl"
	"Runlet/internal/infrastructure/security"
	textdata "Runlet/internal/infrastructure/text_data"
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
		db.Insert(textdata.ClassTable).Rows(goqu.Record{
			"number": "111111",
		}).Executor(),

		db.Insert(textdata.StudentTable).Rows(goqu.Record{
			"name":     "test_student",
			"email":    "test@mail",
			"password": hsh,
			"class_id": 1,
		}).Executor(),

		db.Insert(textdata.CourseTable).Rows(goqu.Record{
			"title":       "test_course",
			"description": "test_description",
		}).Executor(),

		db.Insert(textdata.TeacherTable).Rows(
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

		db.Insert(textdata.ProblemTable).Rows(goqu.Record{
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

func getTestServer(db *goqu.Database) *gin.Engine {
	studentRepo := repositoryimpl.NewStudentRepository(db)
	classRepo := repositoryimpl.NewClassRepository(db)
	courseRepo := repositoryimpl.NewCourseRepository(db)
	teacherRepo := repositoryimpl.NewTeacherRepository(db)
	problemRepo := repositoryimpl.NewProblemRepository(db)
	attemptRepo := repositoryimpl.NewAttemptRepository(db)

	authService := service.NewAuthService(studentRepo, teacherRepo, classRepo)
	studentService := service.NewStudentService(courseRepo, problemRepo, attemptRepo)

	r := gin.New()
	gin.SetMode(gin.TestMode)
	testGroup := r.Group("/test")
	handlers.ConnectStudentHandler(testGroup, authService, studentService)
	handlers.ConnectAuthHandler(testGroup, authService)
	return r
}
