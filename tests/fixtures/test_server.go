package fixtures

import (
	"Runlet/internal/application/service"
	"Runlet/internal/infrastructure/implementations"
	"Runlet/internal/interfaces/http/handlers"
	"net/http/httptest"

	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
)

func GetTestHTTPServer(db *goqu.Database) *httptest.Server {
	studentRepo := implementations.NewStudentRepository(db)
	classRepo := implementations.NewClassRepository(db)
	courseRepo := implementations.NewCourseRepository(db)
	teacherRepo := implementations.NewTeacherRepository(db)
	problemRepo := implementations.NewProblemRepository(db)
	attemptRepo := implementations.NewAttemptRepository(db)

	codeRunner := implementations.NewGRPCRunner()

	authService := service.NewAuthService(studentRepo, teacherRepo, classRepo)
	studentService := service.NewStudentService(courseRepo, problemRepo, attemptRepo, codeRunner)

	r := gin.New()
	gin.SetMode(gin.TestMode)
	testGroup := r.Group("/test")
	handlers.ConnectStudentHandler(testGroup, authService, studentService)
	handlers.ConnectAuthHandler(testGroup, authService)

	server := httptest.NewServer(r)
	return server
}
