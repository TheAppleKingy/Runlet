package service

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/infrastructure/security"
	"context"
	"errors"
)

type StudentAuthService struct {
	studentRepository repository.StudentRepositoryInterface
	classRepository   repository.ClassRepositoryInterface
}

func (s StudentAuthService) Login(ctx context.Context, loginDTO dto.LoginDTO) error {
	student, err := s.studentRepository.GetStudentByEmail(ctx, loginDTO.Email)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.New("no student with provided email")
		}
		return errors.New("Login failed")
	}
	if !security.CheckPassword(loginDTO.Password, student.Password) {
		return errors.New("wrong password")
	}
	return nil
}

func (s StudentAuthService) Register(ctx context.Context, registerDTO dto.RegistrationDTO) error {
	hashedPas, err := security.HashPassword(registerDTO.Password)
	if err != nil {
		return errors.New("fail password")
	}
	class, err := s.classRepository.GetClass(ctx, registerDTO.ClassNum)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.New("no class with provided number")
		}
		return errors.New("error class found")
	}
	_, err = s.studentRepository.CreateStudent(ctx, registerDTO.Name, registerDTO.Email, hashedPas, class.ID)
	if err != nil {
		if ent.IsConstraintError(err) {
			return errors.New("student with this email already exists")
		}
		return errors.New("registration failed")
	}
	return nil
}

type StudentCourseService struct {
	courseRepository repository.CourseRepositoryInterface
}

func (s StudentCourseService) GetStudentCourses(ctx context.Context, studentId int) ([]*ent.Course, error) {
	courses, err := s.courseRepository.GetAllStudentCourses(ctx, studentId)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("no courses found")
		}
		return nil, errors.New("error find courses")
	}
	return courses, nil
}
