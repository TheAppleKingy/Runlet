package service

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/security"
	"Runlet/internal/infrastructure/security/token"
	"context"
	"fmt"
)

type TeacherAuthService struct {
	TeacherRepository repository.TeacherRepositoryInterface
}

func (ts TeacherAuthService) Login(ctx context.Context, loginDTO dto.LoginDTO) (string, error) {
	teacher, err := ts.TeacherRepository.GetTeacherByEmail(ctx, loginDTO.Email)
	if err != nil {
		return "", fmt.Errorf("unable to found teacher: %v", err)

	}
	if !security.CheckPassword(loginDTO.Password, teacher.Password) {
		return "", fmt.Errorf("error password validating: %v", err)
	}
	tokenString, err := token.GetTokenForTeacher(teacher.ID)
	if err != nil {
		return "", fmt.Errorf("login failed: %v", err)
	}
	return tokenString, nil
}

func (ts TeacherAuthService) Register(ctx context.Context, registerDTO dto.RegistrationDTO) error {
	hashedPas, err := security.HashPassword(registerDTO.Password)
	if err != nil {
		return fmt.Errorf("error processing password: %v", err)
	}
	_, err = ts.TeacherRepository.CreateTeacher(ctx, registerDTO.Name, registerDTO.Email, hashedPas)
	if err != nil {
		return fmt.Errorf("unable to create teacher: %v", err)
	}
	return nil
}
