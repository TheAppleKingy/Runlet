package service

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/security"
	"Runlet/internal/infrastructure/security/token"
	"context"
	"fmt"
)

type AuthService struct {
	StudentRepository repository.StudentRepositoryInterface
	TeacherRepository repository.TeacherRepositoryInterface
	ClassRepository   repository.ClassRepositoryInterface
}

func NewAuthService(studentRepo repository.StudentRepositoryInterface, teacherRepo repository.TeacherRepositoryInterface, classRepo repository.ClassRepositoryInterface) *AuthService {
	return &AuthService{
		StudentRepository: studentRepo,
		TeacherRepository: teacherRepo,
		ClassRepository:   classRepo,
	}
}

func (s AuthService) loginStudent(ctx context.Context, email string, password string) (string, error) {
	student, err := s.StudentRepository.GetStudentByEmail(ctx, email)
	if err != nil || student.ID == 0 {
		return "", fmt.Errorf("unable to found student: %v", err)
	}
	if !security.CheckPassword(password, student.Password) {
		return "", fmt.Errorf("wrong password")
	}
	token, err := token.GetTokenForStudent(student.ID)
	if err != nil {
		return "", fmt.Errorf("unable to create token: %v", err)
	}
	return token, nil
}

func (s AuthService) loginTeacher(ctx context.Context, email string, password string) (string, error) {
	teacher, err := s.TeacherRepository.GetTeacherByEmail(ctx, email)
	if err != nil || teacher.ID == 0 {
		return "", fmt.Errorf("unable to found teacher: %v", err)
	}
	if !security.CheckPassword(password, teacher.Password) {
		return "", fmt.Errorf("wrong password")
	}
	token, err := token.GetTokenForTeacher(teacher.ID)
	if err != nil {
		return "", fmt.Errorf("unable to create token: %v", err)
	}
	return token, nil
}

func (s AuthService) RegisterStudent(ctx context.Context, data dto.StudentRegistration) error {
	hashedPas, err := security.HashPassword(data.Password)
	if err != nil {
		return fmt.Errorf("error processing password: %v", err)
	}
	class, err := s.ClassRepository.GetClass(ctx, data.ClassNum)
	if err != nil || class.ID == 0 {
		return fmt.Errorf("unable to found student class: %v", err)
	}
	_, err = s.StudentRepository.CreateStudent(ctx, data.Name, data.Email, hashedPas, class.ID)
	if err != nil {
		return fmt.Errorf("unable to create student: %v", err)
	}
	return nil
}

func (s AuthService) RegisterTeacher(ctx context.Context, data dto.TeacherRegistration) error {
	hashedPas, err := security.HashPassword(data.Password)
	if err != nil {
		return fmt.Errorf("error processing password: %v", err)
	}
	_, err = s.TeacherRepository.CreateTeacher(ctx, data.Name, data.Email, hashedPas)
	if err != nil {
		return fmt.Errorf("unable to create teacher: %v", err)
	}
	return nil
}

func (s AuthService) CheckTeacherExists(ctx context.Context, teacherId int) bool {
	if tch, err := s.TeacherRepository.GetTeacher(ctx, teacherId); err != nil || tch.ID == 0 {
		return false
	}
	return true
}

func (s AuthService) CheckStudentExists(ctx context.Context, studentId int) bool {
	if s, err := s.StudentRepository.GetStudent(ctx, studentId); err != nil || s.ID == 0 {
		return false
	}
	return true
}

func (s AuthService) Login(ctx context.Context, loginData dto.Login) (string, error) {
	if loginData.IsStudent {
		return s.loginStudent(ctx, loginData.Email, loginData.Password)
	}
	return s.loginTeacher(ctx, loginData.Email, loginData.Password)
}
