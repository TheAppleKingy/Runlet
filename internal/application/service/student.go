package service

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/security"
	"Runlet/internal/infrastructure/security/token"
	"context"
	"fmt"
)

type StudentAuthService struct {
	StudentRepository repository.StudentRepositoryInterface
	ClassRepository   repository.ClassRepositoryInterface
}

func NewStudentAuthService(studentRepo repository.StudentRepositoryInterface, classRepo repository.ClassRepositoryInterface) *StudentAuthService {
	return &StudentAuthService{
		StudentRepository: studentRepo,
		ClassRepository:   classRepo,
	}
}

func (s StudentAuthService) Login(ctx context.Context, loginDTO dto.LoginDTO) (string, error) {
	student, err := s.StudentRepository.GetStudentByEmail(ctx, loginDTO.Email)
	if err != nil || student.ID == 0 {
		return "", fmt.Errorf("unable to found student: %v", err)
	}
	if !security.CheckPassword(loginDTO.Password, student.Password) {
		return "", fmt.Errorf("wrong password")
	}
	token, err := token.GetTokenForStudent(student.ID)
	if err != nil {
		return "", fmt.Errorf("unable to create token: %v", err)
	}
	return token, nil
}

func (s StudentAuthService) Register(ctx context.Context, registerDTO dto.RegistrationDTO) error {
	hashedPas, err := security.HashPassword(registerDTO.Password)
	if err != nil {
		return fmt.Errorf("error processing password: %v", err)
	}
	class, err := s.ClassRepository.GetClass(ctx, registerDTO.ClassNum)
	if err != nil || class.ID == 0 {
		return fmt.Errorf("unable to found student class: %v", err)
	}
	_, err = s.StudentRepository.CreateStudent(ctx, registerDTO.Name, registerDTO.Email, hashedPas, class.ID)
	if err != nil {
		return fmt.Errorf("unable to create student: %v", err)
	}
	return nil
}

type StudentService struct {
	CourseRepository repository.CourseRepositoryInterface
}

func NewStudentService(courseRepo repository.CourseRepositoryInterface) *StudentService {
	return &StudentService{
		CourseRepository: courseRepo,
	}
}
