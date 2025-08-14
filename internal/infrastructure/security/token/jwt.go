package token

import (
	"errors"
	"fmt"
	"maps"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getToken(payload map[string]any) (string, error) {
	expireSeconds, err := strconv.Atoi(os.Getenv("JWT_TOKEN_EXPIRE_TIME"))
	if err != nil {
		return "", fmt.Errorf("error getting token exp time: %w", err)
	}
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Second * time.Duration(expireSeconds)).Unix(),
	}
	maps.Copy(claims, payload)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", fmt.Errorf("error get sign token: %w", err)
	}
	return tokenString, nil
}

func GetTokenForStudent(studentId int) (string, error) {
	return getToken(map[string]any{"student_id": studentId})
}

func GetTokenForTeacher(teacherId int) (string, error) {
	return getToken(map[string]any{"teacher_id": teacherId})
}

func getPayloadFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		customError := ErrInvalidToken
		if errors.Is(err, jwt.ErrTokenExpired) {
			customError = ErrTokenExpired
		}
		return nil, customError
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}
	return claims, nil
}

func validateToken(tokenString string, fieldsRequired []string) (jwt.MapClaims, error) {
	payload, err := getPayloadFromToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("error validate token:  %w", err)
	}
	for _, key := range fieldsRequired {
		if _, ok := payload[key]; !ok {
			return nil, fmt.Errorf("error validate token: %w", ErrValidateToken{key})
		}
	}
	return payload, nil
}

func GetStudentFromToken(tokenString string) (int, error) {
	payload, err := validateToken(tokenString, []string{"student_id"})
	if err != nil {
		return 0, fmt.Errorf("error handle token: %w", err)
	}
	studentId, ok := payload["student_id"].(float64)
	if !ok {
		return 0, ErrTokenDataFormat
	}
	return int(studentId), nil
}

func GetTeacherFromToken(tokenString string) (int, error) {
	payload, err := validateToken(tokenString, []string{"teacher_id"})
	if err != nil {
		return 0, fmt.Errorf("error handle token: %w", err)
	}
	studentId, ok := payload["teacher_id"].(float64)
	if !ok {
		return 0, ErrTokenDataFormat
	}
	return int(studentId), nil

}
