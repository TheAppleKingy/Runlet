package handlers

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/application/service"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	StudentService     *service.StudentService
	StudentAuthService *service.StudentAuthService
}

func ConnectStudentHandler(router *gin.RouterGroup, studentService *service.StudentService, studentAuthService *service.StudentAuthService) {
	handler := &StudentHandler{
		StudentService:     studentService,
		StudentAuthService: studentAuthService,
	}
	router.POST("/login", handler.Login)
	router.POST("/logout", handler.Logout)
	router.POST("/registration", handler.Register)
}

// StudentLogin godoc
// @Summary StudentLogin
// @Description Login endpoint for student
// @Tags student/auth
// @Accept  json
// @Produce  json
// @Param loginData body dto.LoginDTO true "Data for login student"
// @Success 200 {object} map[string]string "logged in"
// @Failure 400 {object} map[string]string "logget out"
// @Router /api/student/login [post]
func (h StudentHandler) Login(ctx *gin.Context) {
	var data dto.LoginDTO
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	if _, err := ctx.Cookie("token"); !errors.Is(err, http.ErrNoCookie) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "already logged in",
		})
		return
	}
	tokenString, err := h.StudentAuthService.Login(ctx.Request.Context(), data)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	cookieExpireSeconds, err := strconv.Atoi(os.Getenv("JWT_TOKEN_EXPIRE_TIME"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error": "undefined token expire time",
		})
		return
	}
	ctx.SetCookie(
		"token",
		tokenString,
		cookieExpireSeconds,
		"",
		"",
		false,
		true,
	)
	ctx.JSON(http.StatusOK, gin.H{
		"detail": "logged in",
	})
}

// StudentLogout godoc
// @Summary Logout
// @Description Logout endpoint for student
// @Tags student/auth
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string "logged out"
// @Failure 401 {object} map[string]string
// @Router /api/student/logout [post]
func (h StudentHandler) Logout(ctx *gin.Context) {
	ctx.SetCookie(
		"token",
		"",
		-1,
		"",
		"",
		false,
		true,
	)
	ctx.JSON(http.StatusOK, gin.H{
		"detail": "logged out",
	})
}

// StudentRegister godoc
// @Summary StudentRegister
// @Description Registration endpoint for student
// @Tags student/auth
// @Accept  json
// @Produce  json
// @Param registrationData body dto.RegistrationDTO true "Data for registration student"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/student/registration [post]
func (h StudentHandler) Register(ctx *gin.Context) {
	var data dto.RegistrationDTO
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
	}
	status := http.StatusBadRequest
	if err := h.StudentAuthService.Register(ctx.Request.Context(), data); err != nil {
		if strings.Contains(err.Error(), "registration failed") {
			status = http.StatusInternalServerError
		}
		ctx.AbortWithStatusJSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"detail": "registration successfully",
	})
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
// func (h StudentHandler) GetMyCourses(ctx *gin.Context) {
// 	studentId := ctx.GetInt("student_id")
// 	courses, err := h.StudentService.GetStudentCourses(ctx.Request.Context(), studentId)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, dto.GetCoursesForStudentViewDTO(courses))
// }

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
