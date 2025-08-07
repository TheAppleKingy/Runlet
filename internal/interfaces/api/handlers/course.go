package handlers

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	courseService *service.CourseService
}

func (ch *CourseHandler) GetCourses(ctx *gin.Context) {
	courses, err := ch.courseService.GetCourses(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, courses)
}

func (ch *CourseHandler) CreateCourse(ctx *gin.Context) {
	var createData dto.CourseCreate
	if err := ctx.ShouldBindJSON(&createData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	created, err := ch.courseService.CreateCourse(ctx.Request.Context(), createData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, created)
}
