package student

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/application/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type StudentAuthHandler struct {
	studentAuthService service.StudentAuthService
}

// StudentLogin godoc
// @Summary StudentLogin
// @Description Login endpoint for student
// @Tags student/auth
// @Accept  json
// @Produce  json
// @Param loginData body dto.LoginDTO true "Data for login student"
// @Success 200 {object} map[string]string "logged in"
// @Failure 400 {object} map[string]string
// @Router /api/student/login [post]
func (h StudentAuthHandler) Login(ctx *gin.Context) {
	var data dto.LoginDTO
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
	}
	if err := h.studentAuthService.Login(ctx.Request.Context(), data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"detail": "logged in",
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
// @Router /api/student/registration [post]
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
