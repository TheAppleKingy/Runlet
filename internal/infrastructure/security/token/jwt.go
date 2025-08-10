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

func GetToken(payload map[string]any) (string, error) {
	expireSeconds, err := strconv.Atoi(os.Getenv("JWT_TOKEN_EXPIRE_TIME"))
	if err != nil {
		return "", fmt.Errorf("cannot read JWT_TOKEN_EXPIRE_TIME from env")
	}
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Second * time.Duration(expireSeconds)).Unix(),
	}
	maps.Copy(claims, payload)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(os.Getenv("SECRET_KEY"))
}

func getPayloadFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, fmt.Errorf("invalid sign method for token")
		}
		return os.Getenv("SECRET_KEY"), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

func validateToken(tokenString string, fieldsRequired []string) (jwt.MapClaims, error) {
	payload, err := getPayloadFromToken(tokenString)
	if err != nil {
		return nil, err
	}
	for _, key := range fieldsRequired {
		if _, ok := payload[key]; !ok {
			return nil, fmt.Errorf("payload does not contain field %s", key)
		}
	}
	return payload, nil
}

func GetStudentFromToken(tokenString string) (int, error) {
	payload, err := validateToken(tokenString, []string{"student_id"})
	if err != nil {
		return 0, err
	}
	studentId, ok := payload["student_id"].(int)
	if !ok {
		return 0, errors.New("incorrect token payload data format")
	}
	return studentId, nil
}

func GetTeacherFromToken(tokenString string) (int, error) {
	payload, err := validateToken(tokenString, []string{"teacher_id"})
	if err != nil {
		return 0, err
	}
	studentId, ok := payload["teacher_id"].(int)
	if !ok {
		return 0, errors.New("incorrect token payload data format")
	}
	return studentId, nil

}
