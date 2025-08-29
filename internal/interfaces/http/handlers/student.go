package handlers

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/application/service"
	"Runlet/internal/interfaces/http/middlewares"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	StudentService *service.StudentService
}

func ConnectStudentHandler(parentRouter *gin.RouterGroup, authService *service.AuthService, studentService *service.StudentService) {
	handler := &StudentHandler{
		StudentService: studentService,
	}
	studentRouter := parentRouter.Group("/student", middlewares.AuthMiddleware(authService))

	courseRouter := studentRouter.Group("/courses")
	courseRouter.GET("/", handler.GetMyCourses)
	courseRouter.GET("/:course_id/problems", handler.GetMyProblems)

	problemRouter := studentRouter.Group("/problems")
	problemRouter.POST("/:problem_id/send_solution", handler.SendSolution)
}

// GetCourses godoc
// @Summary GetMyCourses
// @Description Returns all student's courses
// @Tags student/courses
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.CourseForStudent "Example of course data"
// @Failure 400 {object} map[string]string
// @Router /api/student/courses/ [get]
func (h StudentHandler) GetMyCourses(ctx *gin.Context) {
	studentId := ctx.GetInt("student_id")
	courses, err := h.StudentService.GetStudentCourses(ctx.Request.Context(), studentId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, courses)
}

// GetMyProblems godoc
// @Summary GetMyProblems
// @Description Returns student's problems by course
// @Tags student/courses
// @Accept  json
// @Produce  json
// @Param course_id path int true "Course ID"
// @Success 200 {array} entities.Problem "Example of problem data"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/student/courses/{course_id}/problems [get]
func (h StudentHandler) GetMyProblems(ctx *gin.Context) {
	studentId := ctx.GetInt("student_id")
	courseId, err := strconv.Atoi(ctx.Param("course_id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid course id"})
		return
	}
	problems, err := h.StudentService.GetStudentProblems(ctx.Request.Context(), studentId, courseId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, problems)
}

// SendSolution godoc
// @Summary SendSolution
// @Description Sending solutions code
// @Tags student/problems
// @Accept  json
// @Produce  json
// @Param problem_id path int true "Problem ID"
// @Param sendCodeData body dto.CodeSolution true "Data for sending code solution"
// @Success 200 {array} entities.Problem "Example of response"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/student/problems/{problem_id}/send_solution [post]
func (h StudentHandler) SendSolution(ctx *gin.Context) {
	studentId := ctx.GetInt("student_id")
	problemId, err := strconv.Atoi(ctx.Param("problem_id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid problem id"})
		return
	}
	var data dto.CodeSolution
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	go h.StudentService.SendCodeSolution(context.Background(), studentId, problemId, data)
	ctx.JSON(http.StatusOK, gin.H{
		"detail": "solution sent to server, wait for results",
	})
}
