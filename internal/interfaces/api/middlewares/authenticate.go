package middlewares

import (
	"Runlet/internal/infrastructure/security/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Cookie("token")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token was not provide",
			})
		}
		teacherId, err := token.GetTeacherFromToken(tokenString)
		if err != nil {
			studentId, err := token.GetStudentFromToken(tokenString)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
			}
			ctx.Set("student_id", studentId)
			ctx.Next()
		}
		ctx.Set("teacher_id", teacherId)
		ctx.Next()
	}
}
