package integration

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/application/service"
	"Runlet/internal/infrastructure/security/token"
	"Runlet/internal/pkg/errs"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const contentType string = "application/json"

func TestStudentLoginOk(t *testing.T) {
	body, _ := json.Marshal(dto.Login{
		Email:     "test@mail",
		Password:  "test_password",
		IsStudent: true,
	})

	resp, err := http.Post(MainURL+"/auth/login", contentType, bytes.NewBuffer(body))
	assert.NoError(t, err)
	defer resp.Body.Close()
	expected, _ := json.Marshal(map[string]string{"detail": "logged in"})
	recieved, _ := io.ReadAll(resp.Body)
	cookies := resp.Cookies()
	assert.Equal(t, len(cookies), 1)
	assert.Equal(t, cookies[0].Name, "token")
	assert.True(t, cookies[0].MaxAge > 0)
	student, _ := token.GetStudentFromToken(cookies[0].Value)
	assert.Equal(t, student, 1)
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, string(expected), string(recieved))
}

func TestStudentLoginNotExists(t *testing.T) {
	body, _ := json.Marshal(dto.Login{
		Email:     "est@mail",
		Password:  "test_password",
		IsStudent: true,
	})

	resp, err := http.Post(MainURL+"/auth/login", contentType, bytes.NewBuffer(body))
	assert.NoError(t, err)
	defer resp.Body.Close()
	expected, err := json.Marshal(map[string]string{"error": errs.ConcatErrors(service.ErrStudentNotFound, nil).Error()})
	assert.NoError(t, err)
	recieved, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 400)
	assert.Equal(t, string(expected), string(recieved))
}

func TestStudentLoginWrongPass(t *testing.T) {
	body, _ := json.Marshal(dto.Login{
		Email:     "test@mail",
		Password:  "test_pasword",
		IsStudent: true,
	})

	resp, err := http.Post(MainURL+"/auth/login", contentType, bytes.NewBuffer(body))
	assert.NoError(t, err)
	defer resp.Body.Close()
	expected, err := json.Marshal(map[string]string{"error": "wrong password"})
	assert.NoError(t, err)
	recieved, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 400)
	assert.Equal(t, string(expected), string(recieved))
}

func TestStudentRegistrationOk(t *testing.T) {
	t.Cleanup(func() {
		DB.Exec("delete from students where email = $1", "new@mail")
	})
	body, _ := json.Marshal(dto.StudentRegistration{
		Name:     "new_name",
		Email:    "new@mail",
		Password: "test_password",
		ClassNum: "111111",
	})

	resp, err := http.Post(MainURL+"/auth/registration_student", contentType, bytes.NewBuffer(body))
	assert.NoError(t, err)
	defer resp.Body.Close()
	expected, _ := json.Marshal(map[string]string{"detail": "registration successfully"})
	recieved, _ := io.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, string(expected), string(recieved))
}

func TestStudentRegistrationEmailAlreadyExists(t *testing.T) {
	body, _ := json.Marshal(dto.StudentRegistration{
		Name:     "new_name",
		Email:    "test@mail",
		Password: "test_password",
		ClassNum: "111111",
	})

	resp, err := http.Post(MainURL+"/auth/registration_student", contentType, bytes.NewBuffer(body))
	assert.NoError(t, err)
	defer resp.Body.Close()
	expected, _ := json.Marshal(map[string]string{"error": errs.ConcatErrors(service.ErrRegisterStudent, errors.New("pq: duplicate key value violates unique constraint \"students_email_key\"")).Error()})
	recieved, _ := io.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, 400)
	assert.Equal(t, string(expected), string(recieved))
}

func TestStudentRegistrationNoClassExists(t *testing.T) {
	body, _ := json.Marshal(dto.StudentRegistration{
		Name:     "new_name",
		Email:    "new@mail",
		Password: "test_password",
		ClassNum: "111011",
	})

	resp, err := http.Post(MainURL+"/auth/registration_student", contentType, bytes.NewBuffer(body))
	assert.NoError(t, err)
	defer resp.Body.Close()
	expected, _ := json.Marshal(map[string]string{"error": errs.ConcatErrors(service.ErrClassNotFound, nil).Error()})
	recieved, _ := io.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, 400)
	assert.Equal(t, string(expected), string(recieved))
}

func TestStudentRegistrationInvalidEmailFormat(t *testing.T) {
	body, _ := json.Marshal(dto.StudentRegistration{
		Name:     "new_name",
		Email:    "newmail",
		Password: "test_password",
		ClassNum: "111111",
	})

	resp, err := http.Post(MainURL+"/auth/registration_student", contentType, bytes.NewBuffer(body))
	assert.NoError(t, err)
	defer resp.Body.Close()
	expected, _ := json.Marshal(map[string]string{"error": errs.ConcatErrors(service.ErrRegisterStudent, errors.New("pq: value for domain email_type violates check constraint \"email_type_check\"")).Error()})
	recieved, _ := io.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, 400)
	assert.Equal(t, string(expected), string(recieved))
}

func TestLogout(t *testing.T) {
	body, _ := json.Marshal(nil)
	resp, err := http.Post(MainURL+"/auth/logout", contentType, bytes.NewBuffer(body))
	assert.NoError(t, err)
	defer resp.Body.Close()
	expected, _ := json.Marshal(map[string]string{"detail": "logged out"})
	recieved, _ := io.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, string(expected), string(recieved))
	cookies := resp.Cookies()
	assert.Equal(t, len(cookies), 1)
	assert.Equal(t, cookies[0].Name, "token")
	assert.Equal(t, cookies[0].MaxAge, -1)
}

func TestTeacherLoginOk(t *testing.T) {
	body, _ := json.Marshal(dto.Login{
		Email:     "test_t@mail",
		Password:  "test_password",
		IsStudent: false,
	})

	resp, err := http.Post(MainURL+"/auth/login", contentType, bytes.NewBuffer(body))
	assert.NoError(t, err)
	defer resp.Body.Close()
	expected, _ := json.Marshal(map[string]string{"detail": "logged in"})
	recieved, _ := io.ReadAll(resp.Body)
	cookies := resp.Cookies()
	assert.Equal(t, len(cookies), 1)
	assert.Equal(t, cookies[0].Name, "token")
	assert.True(t, cookies[0].MaxAge > 0)
	teacher, _ := token.GetTeacherFromToken(cookies[0].Value)
	assert.Equal(t, teacher, 1)
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, string(expected), string(recieved))
}

func TestTeacherLoginNotExists(t *testing.T) {
	body, _ := json.Marshal(dto.Login{
		Email:     "est@mail",
		Password:  "test_password",
		IsStudent: false,
	})

	resp, err := http.Post(MainURL+"/auth/login", contentType, bytes.NewBuffer(body))
	assert.NoError(t, err)
	defer resp.Body.Close()
	expected, err := json.Marshal(map[string]string{"error": errs.ConcatErrors(service.ErrTeacherNotFound, nil).Error()})
	assert.NoError(t, err)
	recieved, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 400)
	assert.Equal(t, string(recieved), string(expected))
}

func TestTeacherLoginWrongPass(t *testing.T) {
	body, _ := json.Marshal(dto.Login{
		Email:     "test_t@mail",
		Password:  "test_pasword",
		IsStudent: false,
	})

	resp, err := http.Post(MainURL+"/auth/login", contentType, bytes.NewBuffer(body))
	assert.NoError(t, err)
	defer resp.Body.Close()
	expected, err := json.Marshal(map[string]string{"error": service.ErrWrongPassword.Error()})
	assert.NoError(t, err)
	recieved, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 400)
	assert.Equal(t, string(recieved), string(expected))
}

func TestTeacherRegistrationOk(t *testing.T) {
	t.Cleanup(func() {
		DB.Exec("delete from students where email = $1", "new@mail")
	})
	body, _ := json.Marshal(dto.TeacherRegistration{
		Name:     "new_name",
		Email:    "new@mail",
		Password: "test_password",
	})

	resp, err := http.Post(MainURL+"/auth/registration_teacher", contentType, bytes.NewBuffer(body))
	assert.NoError(t, err)
	defer resp.Body.Close()
	expected, _ := json.Marshal(map[string]string{"detail": "registration successfully"})
	recieved, _ := io.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, string(expected), string(recieved))
}

func TestTeacherRegistrationEmailAlreadyExists(t *testing.T) {
	body, _ := json.Marshal(dto.TeacherRegistration{
		Name:     "new_name",
		Email:    "test_t@mail",
		Password: "test_password",
	})

	resp, err := http.Post(MainURL+"/auth/registration_teacher", contentType, bytes.NewBuffer(body))
	assert.NoError(t, err)
	defer resp.Body.Close()
	expected, _ := json.Marshal(map[string]string{"error": errs.ConcatErrors(service.ErrRegisterTeacher, errors.New("pq: duplicate key value violates unique constraint \"teachers_email_key\"")).Error()})
	recieved, _ := io.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, 400)
	assert.Equal(t, string(expected), string(recieved))
}

func TestTeacherRegistrationInvalidEmailFormat(t *testing.T) {
	body, _ := json.Marshal(dto.TeacherRegistration{
		Name:     "new_name",
		Email:    "newmail",
		Password: "test_password",
	})

	resp, err := http.Post(MainURL+"/auth/registration_teacher", contentType, bytes.NewBuffer(body))
	assert.NoError(t, err)
	defer resp.Body.Close()
	expected, _ := json.Marshal(map[string]string{"error": errs.ConcatErrors(service.ErrRegisterTeacher, errors.New("pq: value for domain email_type violates check constraint \"email_type_check\"")).Error()})
	recieved, _ := io.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, 400)
	assert.Equal(t, string(expected), string(recieved))
}
