package token

import (
	"Runlet/internal/infrastructure/config"
	"Runlet/internal/pkg/errs"
	"maps"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getToken(payload map[string]any) (string, error) {
	expireSeconds := config.AuthConfig.TokenExpireTime
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Second * time.Duration(expireSeconds)).Unix(),
	}
	maps.Copy(claims, payload)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signKey := config.AuthConfig.Secret
	tokenString, err := token.SignedString([]byte(signKey))
	if err != nil {
		return "", errs.WrapErrors(ErrSigningToken, err)
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
		return []byte(config.AuthConfig.Secret), nil
	})
	if err != nil {
		return nil, errs.WrapErrors(ErrInvalidToken, err)
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
		return nil, errs.WrapErrors(ErrValidateToken, err)
	}
	for _, key := range fieldsRequired {
		if _, ok := payload[key]; !ok {
			return nil, errs.WrapErrors(ErrValidateToken, ErrTokenNoRequiredData{key})
		}
	}
	return payload, nil
}

func GetStudentFromToken(tokenString string) (int, error) {
	payload, err := validateToken(tokenString, []string{"student_id"})
	if err != nil {
		return 0, err
	}
	studentId, ok := payload["student_id"].(float64)
	if !ok {
		return 0, errs.WrapErrors(ErrValidateToken, ErrTokenDataFormat)
	}
	return int(studentId), nil
}

func GetTeacherFromToken(tokenString string) (int, error) {
	payload, err := validateToken(tokenString, []string{"teacher_id"})
	if err != nil {
		return 0, err
	}
	studentId, ok := payload["teacher_id"].(float64)
	if !ok {
		return 0, errs.WrapErrors(ErrValidateToken, ErrTokenDataFormat)
	}
	return int(studentId), nil

}
