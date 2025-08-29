package middlewares

import (
	"Runlet/internal/application/service"
	"Runlet/internal/infrastructure/security/token"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService *service.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("token")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token was not provide",
			})
			return
		}
		teacherId, err := token.GetTeacherFromToken(tokenString)
		if err != nil {
			studentId, err := token.GetStudentFromToken(tokenString)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				return
			}
			if !authService.CheckStudentExists(context.Background(), studentId) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "unable to find student",
				})
				return
			}
			ctx.Set("student_id", studentId)
			ctx.Next()
			return
		}
		if !authService.CheckTeacherExists(context.Background(), teacherId) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unable to find teacher",
			})
			return
		}
		ctx.Set("teacher_id", teacherId)
		ctx.Next()
	}
}
