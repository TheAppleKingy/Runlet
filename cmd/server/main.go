package main

import (
	"Runlet/internal/application/service"
	"Runlet/internal/infrastructure/repositoryimpl"
	"Runlet/internal/interfaces/http/handlers"
	"database/sql"
	"log/slog"
	"os"

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
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		slog.Error("Database url did not set in environment")
		os.Exit(1)
	}
	dbClient, err := sql.Open("postgres", dbUrl)
	if err != nil {
		slog.Error("error database connection", "error", err)
		os.Exit(1)
	}
	db := goqu.New("postgres", dbClient)
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRouter := router.Group("/api")

	classRepository := repositoryimpl.NewClassRepository(db)
	courseRepository := repositoryimpl.NewCourseRepository(db)
	studentRepository := repositoryimpl.NewStudentRepository(db)
	// teacherRepository := repositoryimpl.NewTeacherRepository(db)
	// problemRepository := repositoryimpl.NewProblemRepository(db)

	studentService := service.NewStudentService(courseRepository)
	studentAuthService := service.NewStudentAuthService(studentRepository, classRepository)

	stRouter := apiRouter.Group("/student")
	handlers.ConnectStudentHandler(stRouter, studentService, studentAuthService)

	//nolint:errcheck
	router.Run(":8080")
}
