package main

import (
	"Runlet/internal/application/service"
	"Runlet/internal/infrastructure/repositoryimpl"
	"Runlet/internal/interfaces/http/handlers"
	"database/sql"
	"fmt"
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
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	if dbName == "" || dbUser == "" || dbPassword == "" {
		slog.Error("No db connection params in env")
		os.Exit(1)
	}
	dbUrl := fmt.Sprintf("postgres://%s:%s@database:5432/%s?sslmode=disable", dbUser, dbPassword, dbName)
	dbClient, err := sql.Open("postgres", dbUrl)
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

	classRepository := repositoryimpl.NewClassRepository(db)
	courseRepository := repositoryimpl.NewCourseRepository(db)
	studentRepository := repositoryimpl.NewStudentRepository(db)
	teacherRepository := repositoryimpl.NewTeacherRepository(db)
	problemRepository := repositoryimpl.NewProblemRepository(db)
	attemptRepository := repositoryimpl.NewAttemptRepository(db)

	studentService := service.NewStudentService(courseRepository, problemRepository, attemptRepository)
	authService := service.NewAuthService(studentRepository, teacherRepository, classRepository)

	handlers.ConnectAuthHandler(apiRouter, authService)
	handlers.ConnectStudentHandler(apiRouter, authService, studentService)

	//nolint:errcheck
	router.Run(":8080")
}
