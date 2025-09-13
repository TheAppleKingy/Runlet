package main

import (
	"Runlet/internal/application/service"
	"Runlet/internal/infrastructure/config"
	"Runlet/internal/infrastructure/implementations"
	"Runlet/internal/interfaces/http/handlers"
	"database/sql"
	"log/slog"
	"os"
	"strconv"

	_ "Runlet/docs"

	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Runlet API
// @version 1.0
// @description API documentation for Runlet
// @host localhost:8081
// @BasePath /
func main() {
	config.LoadConfigs()

	dbClient, err := sql.Open("postgres", config.DBConfig.URL)
	if err != nil {
		slog.Error("error database connection", "error", err)
		os.Exit(1)
	}

	db := goqu.New("postgres", dbClient)

	router := gin.Default()
	if debug, _ := strconv.ParseBool(os.Getenv("DEBUG")); debug {
		router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	apiRouter := router.Group("/api")

	classRepository := implementations.NewClassRepository(db)
	courseRepository := implementations.NewCourseRepository(db)
	studentRepository := implementations.NewStudentRepository(db)
	teacherRepository := implementations.NewTeacherRepository(db)
	problemRepository := implementations.NewProblemRepository(db)
	attemptRepository := implementations.NewAttemptRepository(db)

	codeRunner := implementations.NewGRPCRunner()

	studentService := service.NewStudentService(courseRepository, problemRepository, attemptRepository, codeRunner)
	authService := service.NewAuthService(studentRepository, teacherRepository, classRepository)

	handlers.ConnectAuthHandler(apiRouter, authService)
	handlers.ConnectStudentHandler(apiRouter, authService, studentService)

	//nolint:errcheck
	router.Run(":8080")
}
