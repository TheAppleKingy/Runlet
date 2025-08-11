package student

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/infrastructure/security"
	"Runlet/internal/infrastructure/security/token"
	"context"
	"errors"
)

type StudentAuthService struct {
	StudentRepository repository.StudentRepositoryInterface
	ClassRepository   repository.ClassRepositoryInterface
}

func (s StudentAuthService) Login(ctx context.Context, loginDTO dto.LoginDTO) (string, error) {
	student, err := s.StudentRepository.GetStudentByEmail(ctx, loginDTO.Email)
	if err != nil {
		if ent.IsNotFound(err) {
			return "", errors.New("no student with provided email")
		}
		return "", errors.New("login failed")
	}
	if !security.CheckPassword(loginDTO.Password, student.Password) {
		return "", errors.New("wrong password")
	}
	token, err := token.GetTokenForStudent(student.ID)
	if err != nil {
		return "", errors.New("cannot make authentication token")
	}
	return token, nil
}

func (s StudentAuthService) Register(ctx context.Context, registerDTO dto.RegistrationDTO) error {
	hashedPas, err := security.HashPassword(registerDTO.Password)
	if err != nil {
		return errors.New("fail password")
	}
	class, err := s.ClassRepository.GetClass(ctx, registerDTO.ClassNum)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.New("no class with provided number")
		}
		return errors.New("error class found")
	}
	_, err = s.StudentRepository.CreateStudent(ctx, registerDTO.Name, registerDTO.Email, hashedPas, class.ID)
	if err != nil {
		if ent.IsConstraintError(err) {
			return errors.New("student with this email already exists")
		}
		return errors.New("registration failed")
	}
	return nil
}
