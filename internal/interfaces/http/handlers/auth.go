package handlers

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/application/service"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func ConnectAuthHandler(parentRouter *gin.RouterGroup, authService *service.AuthService) {
	handler := AuthHandler{
		AuthService: authService,
	}
	authRouter := parentRouter.Group("/auth")
	authRouter.POST("/login", handler.Login)
	authRouter.POST("/logout", handler.Logout)
	authRouter.POST("/registration_student", handler.RegisterStudent)
	authRouter.POST("/registration_teacher", handler.RegisterTeacher)

}

// Login godoc
// @Summary Login
// @Description Login endpoint
// @Tags auth
// @Accept  json
// @Produce  json
// @Param loginData body dto.Login true "Data for login"
// @Success 200 {object} map[string]string "logged in"
// @Failure 400 {object} map[string]string "logget out"
// @Router /api/auth/login [post]
func (h AuthHandler) Login(ctx *gin.Context) {
	var data dto.Login
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
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

	token, err := h.AuthService.Login(ctx.Request.Context(), data)
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
		token,
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

// Logout godoc
// @Summary Logout
// @Description Logout endpoint
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string "logged out"
// @Failure 401 {object} map[string]string
// @Router /api/auth/logout [post]
func (h AuthHandler) Logout(ctx *gin.Context) {
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
// @Tags auth
// @Accept  json
// @Produce  json
// @Param registrationData body dto.StudentRegistration true "Data for registration student"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/auth/registration_student [post]
func (h AuthHandler) RegisterStudent(ctx *gin.Context) {
	var data dto.StudentRegistration
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
	}
	status := http.StatusBadRequest
	if err := h.AuthService.RegisterStudent(ctx.Request.Context(), data); err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"detail": "registration successfully",
	})
}

// TeacherRegister godoc
// @Summary TeacherRegister
// @Description Registration endpoint for teacher
// @Tags auth
// @Accept  json
// @Produce  json
// @Param registrationData body dto.TeacherRegistration true "Data for registration teacher"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/auth/registration_teacher [post]
func (h AuthHandler) RegisterTeacher(ctx *gin.Context) {
	var data dto.TeacherRegistration
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
	}
	status := http.StatusBadRequest
	if err := h.AuthService.RegisterTeacher(ctx.Request.Context(), data); err != nil {
		ctx.AbortWithStatusJSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"detail": "registration successfully",
	})
}
