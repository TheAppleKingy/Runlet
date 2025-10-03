package integration

import (
	"Runlet/internal/application/service"
	"Runlet/internal/infrastructure/config"
	"Runlet/internal/infrastructure/implementations"
	"Runlet/internal/interfaces/http/handlers"
	"log/slog"
	"net/http/httptest"
	"os"

	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
)

func getTestHTTPServer(db *goqu.Database) *httptest.Server {
	if err := config.LoadConfigs(); err != nil {
		slog.Error("cannot load config data", "error", err)
		os.Exit(1)
	}
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
