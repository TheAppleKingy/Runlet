package student

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/application/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type StudentCourseHandler struct {
	studentService *service.StudentCourseService
}

// GetCourses godoc
// @Summary GetMyCourses
// @Description Returns all student's courses
// @Tags student/courses
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.CourseViewDTO
// @Failure 400 {object} map[string]string
// @Router /api/student/my_courses [get]
func (h StudentCourseHandler) GetMyCourses(ctx *gin.Context) {
	studentId := ctx.GetInt("student_id")
	courses, err := h.studentService.GetStudentCourses(ctx.Request.Context(), studentId)
	if err != nil {
		status := http.StatusInternalServerError
		if strings.Contains(err.Error(), "no courses found") {
			status = http.StatusOK
		}
		ctx.AbortWithStatusJSON(status, gin.H{
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
