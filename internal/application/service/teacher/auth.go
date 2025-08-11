package teacher

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/domain/repository"
	"Runlet/internal/infrastructure/ent"
	"Runlet/internal/infrastructure/security"
	"context"
	"errors"
)

type TeacherAuthService struct {
	TeacherRepository repository.TeacherRepositoryInterface
}

func (ts TeacherAuthService) Login(ctx context.Context, loginDTO dto.LoginDTO) error {
	teacher, err := ts.TeacherRepository.GetTeacherByEmail(ctx, loginDTO.Email)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.New("no teacher with provided email")
		}
		return errors.New("Login failed")
	}
	if !security.CheckPassword(loginDTO.Password, teacher.Password) {
		return errors.New("wrong password")
	}
	return nil
}

func (ts TeacherAuthService) Register(ctx context.Context, registerDTO dto.RegistrationDTO) error {
	hashedPas, err := security.HashPassword(registerDTO.Password)
	if err != nil {
		return errors.New("fail password proccessing")
	}
	_, err = ts.TeacherRepository.CreateTeacher(ctx, registerDTO.Name, registerDTO.Email, hashedPas)
	if err != nil {
		if ent.IsConstraintError(err) {
			return errors.New("teacher with this email already exists")
		}
		return errors.New("registration failed")
	}
	return nil
}
