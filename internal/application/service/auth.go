package service

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/domain/interfaces"
	"Runlet/internal/infrastructure/security"
	"Runlet/internal/infrastructure/security/token"
	"Runlet/internal/pkg/errs"
	"context"
)

type AuthService struct {
	StudentRepository interfaces.StudentRepository
	TeacherRepository interfaces.TeacherRepository
	ClassRepository   interfaces.ClassRepository
}

func NewAuthService(studentRepo interfaces.StudentRepository, teacherRepo interfaces.TeacherRepository, classRepo interfaces.ClassRepository) *AuthService {
	return &AuthService{
		StudentRepository: studentRepo,
		TeacherRepository: teacherRepo,
		ClassRepository:   classRepo,
	}
}

func (s AuthService) loginStudent(ctx context.Context, email string, password string) (string, error) {
	student, err := s.StudentRepository.GetStudentByEmail(ctx, email)
	if err != nil || student.ID == 0 {
		return "", errs.ConcatErrors(ErrStudentNotFound, err)
	}
	if !security.CheckPassword(password, student.Password) {
		return "", ErrWrongPassword
	}
	token, err := token.GetTokenForStudent(student.ID)
	if err != nil {
		return "", errs.ConcatErrors(ErrCreateToken, err)
	}
	return token, nil
}

func (s AuthService) loginTeacher(ctx context.Context, email string, password string) (string, error) {
	teacher, err := s.TeacherRepository.GetTeacherByEmail(ctx, email)
	if err != nil || teacher.ID == 0 {
		return "", errs.ConcatErrors(ErrTeacherNotFound, err)
	}
	if !security.CheckPassword(password, teacher.Password) {
		return "", ErrWrongPassword
	}
	token, err := token.GetTokenForTeacher(teacher.ID)
	if err != nil {
		return "", errs.ConcatErrors(ErrCreateToken, err)
	}
	return token, nil
}

func (s AuthService) RegisterStudent(ctx context.Context, data dto.StudentRegistration) error {
	hashedPas, err := security.HashPassword(data.Password)
	if err != nil {
		return errs.ConcatErrors(ErrProcessingPassword, err)
	}
	class, err := s.ClassRepository.GetClass(ctx, data.ClassNum)
	if err != nil || class.ID == 0 {
		return errs.ConcatErrors(ErrClassNotFound, err)
	}
	_, err = s.StudentRepository.CreateStudent(ctx, data.Name, data.Email, hashedPas, class.ID)
	if err != nil {
		return errs.ConcatErrors(ErrRegisterStudent, err)
	}
	return nil
}

func (s AuthService) RegisterTeacher(ctx context.Context, data dto.TeacherRegistration) error {
	hashedPas, err := security.HashPassword(data.Password)
	if err != nil {
		return errs.ConcatErrors(ErrProcessingPassword, err)
	}
	_, err = s.TeacherRepository.CreateTeacher(ctx, data.Name, data.Email, hashedPas)
	if err != nil {
		return errs.ConcatErrors(ErrRegisterTeacher, err)
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
