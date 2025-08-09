package student

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentCourseHandler struct {
	CourseService *service.CourseService
}

// GetCourses godoc
// @Summary Get all courses
// @Description Returns all courses from the database
// @Tags courses
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.CourseViewDTO
// @Failure 400 {object} map[string]string
// @Router /api/student/my_courses [get]
func (ch *StudentCourseHandler) GetCourses(ctx *gin.Context) {
	courses, err := ch.CourseService.GetAllCourses(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, dto.GetCoursesForStudentViewDTO(courses))
}

// GetCourses godoc
// @Summary Create course
// @Description Create course and return it
// @Tags courses
// @Accept  json
// @Produce  json
//
//	@Param   createData  body  dto.CourseCreateDTO  true  "Course creation data"  example(`{
//	  "title": "Introduction to Go",
//	  "description": "A beginner-friendly Go course",
//	  "classesIds": [1, 2, 3]
//	}`)
//
// @Success 200 {array} dto.CourseViewDTO
// @Failure 400 {object} map[string]string
// @Router /api/student/create_course [post]
func (ch *StudentCourseHandler) CreateCourse(ctx *gin.Context) {
	var createData dto.CourseCreateDTO
	if err := ctx.ShouldBindJSON(&createData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	created, err := ch.CourseService.CreateCourse(ctx.Request.Context(), createData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.GetCourseViewNoEdgesDTO(created))
}
