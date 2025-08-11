package student

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/application/service/student"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type StudentAuthHandler struct {
	studentAuthService student.StudentAuthService
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
// @Router /api/student/auth/login [post]
func (h StudentAuthHandler) Login(ctx *gin.Context) {
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
	tokenString, err := h.studentAuthService.Login(ctx.Request.Context(), data)
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
// @Router /api/student/auth/logout [post]
func (h StudentAuthHandler) Logout(ctx *gin.Context) {
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
// @Success 200 {object} map[string]string "registration successfully"
// @Failure 400 {object} map[string]string
// @Router /api/student/auth/registration [post]
func (h StudentAuthHandler) Register(ctx *gin.Context) {
	var data dto.RegistrationDTO
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
	}
	status := http.StatusBadRequest
	if err := h.studentAuthService.Register(ctx.Request.Context(), data); err != nil {
		if strings.Contains(err.Error(), "registration failed") {
			status = http.StatusInternalServerError
		}
		ctx.AbortWithStatusJSON(status, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"detail": "registration successfully",
	})
}
