package student

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/application/service/student"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentCourseHandler struct {
	studentService *student.StudentCourseService
}

// GetCourses godoc
// @Summary GetMyCourses
// @Description Returns all student's courses
// @Tags student/courses
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.CourseViewDTO
// @Failure 400 {object} map[string]string
// @Router /api/student/course/my_courses [get]
func (h StudentCourseHandler) GetMyCourses(ctx *gin.Context) {
	studentId := ctx.GetInt("student_id")
	courses, err := h.studentService.GetStudentCourses(ctx.Request.Context(), studentId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
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
