package tests

import (
	"Runlet/internal/application/dto"
	"Runlet/internal/infrastructure/security/token"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

const contentType string = "application/json"

func TestStudentLoginOk(t *testing.T) {
	body, _ := json.Marshal(dto.LoginDTO{
		Email:    "test@mail",
		Password: "test_password",
	})

	resp, err := http.Post(MainURL+"/student/login", contentType, bytes.NewBuffer(body))
	assert.NilError(t, err)
	defer resp.Body.Close()
	expected, _ := json.Marshal(map[string]string{"detail": "logged in"})
	recieved, _ := io.ReadAll(resp.Body)
	cookies := resp.Cookies()
	assert.Equal(t, len(cookies), 1)
	assert.Equal(t, cookies[0].Name, "token")
	assert.Assert(t, cookies[0].MaxAge > 0)
	student, _ := token.GetStudentFromToken(cookies[0].Value)
	assert.Equal(t, student, 1)
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, string(expected), string(recieved))
}

func TestStudentLoginNotExists(t *testing.T) {
	body, _ := json.Marshal(dto.LoginDTO{
		Email:    "est@mail",
		Password: "test_password",
	})

	resp, err := http.Post(MainURL+"/student/login", contentType, bytes.NewBuffer(body))
	assert.NilError(t, err)
	defer resp.Body.Close()
	expected, err := json.Marshal(map[string]string{"error": "unable to found student: <nil>"})
	assert.NilError(t, err)
	recieved, err := io.ReadAll(resp.Body)
	assert.NilError(t, err)
	assert.Equal(t, resp.StatusCode, 400)
	assert.Equal(t, string(recieved), string(expected))
}

func TestStudentLoginWrongPass(t *testing.T) {
	body, _ := json.Marshal(dto.LoginDTO{
		Email:    "test@mail",
		Password: "test_pasword",
	})

	resp, err := http.Post(MainURL+"/student/login", contentType, bytes.NewBuffer(body))
	assert.NilError(t, err)
	defer resp.Body.Close()
	expected, err := json.Marshal(map[string]string{"error": "wrong password"})
	assert.NilError(t, err)
	recieved, err := io.ReadAll(resp.Body)
	assert.NilError(t, err)
	assert.Equal(t, resp.StatusCode, 400)
	assert.Equal(t, string(recieved), string(expected))
}

func TestStudentLoginNoExpTime(t *testing.T) {
	body, _ := json.Marshal(dto.LoginDTO{
		Email:    "test@mail",
		Password: "test_password",
	})
	expTime := os.Getenv("JWT_TOKEN_EXPIRE_TIME")
	os.Unsetenv("JWT_TOKEN_EXPIRE_TIME")
	resp, err := http.Post(MainURL+"/student/login", contentType, bytes.NewBuffer(body))
	os.Setenv("JWT_TOKEN_EXPIRE_TIME", expTime)
	assert.NilError(t, err)
	defer resp.Body.Close()
	expected, err := json.Marshal(map[string]string{"error": "unable to create token: error getting token exp time: strconv.Atoi: parsing \"\": invalid syntax"})
	assert.NilError(t, err)
	recieved, err := io.ReadAll(resp.Body)
	assert.NilError(t, err)
	assert.Equal(t, resp.StatusCode, 400)
	assert.Equal(t, string(recieved), string(expected))
}

func TestStudentLoginNoSecret(t *testing.T) {
	body, _ := json.Marshal(dto.LoginDTO{
		Email:    "test@mail",
		Password: "test_password",
	})
	sKey := os.Getenv("SECRET_KEY")
	os.Unsetenv("SECRET_KEY")
	resp, err := http.Post(MainURL+"/student/login", contentType, bytes.NewBuffer(body))
	os.Setenv("SECRET_KEY", sKey)
	assert.NilError(t, err)
	defer resp.Body.Close()
	expected, err := json.Marshal(map[string]string{"error": "unable to create token: no sign key"})
	assert.NilError(t, err)
	recieved, err := io.ReadAll(resp.Body)
	assert.NilError(t, err)
	assert.Equal(t, resp.StatusCode, 400)
	assert.Equal(t, string(recieved), string(expected))
}

func TestRegistrationOk(t *testing.T) {
	t.Cleanup(func() {
		db.Exec("delete from students where email = $1", "new@mail")
	})
	body, _ := json.Marshal(dto.RegistrationDTO{
		Name:     "new_name",
		Email:    "new@mail",
		Password: "test_password",
		ClassNum: "111111",
	})

	resp, err := http.Post(MainURL+"/student/registration", contentType, bytes.NewBuffer(body))
	assert.NilError(t, err)
	defer resp.Body.Close()
	expected, _ := json.Marshal(map[string]string{"detail": "registration successfully"})
	recieved, _ := io.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, string(expected), string(recieved))
}

func TestRegistrationEmailAlreadyExists(t *testing.T) {
	body, _ := json.Marshal(dto.RegistrationDTO{
		Name:     "new_name",
		Email:    "test@mail",
		Password: "test_password",
		ClassNum: "111111",
	})

	resp, err := http.Post(MainURL+"/student/registration", contentType, bytes.NewBuffer(body))
	assert.NilError(t, err)
	defer resp.Body.Close()
	expected, _ := json.Marshal(map[string]string{"error": "unable to create student: pq: duplicate key value violates unique constraint \"students_email_key\""})
	recieved, _ := io.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, 400)
	assert.Equal(t, string(expected), string(recieved))
}

func TestRegistrationNoClassExists(t *testing.T) {
	body, _ := json.Marshal(dto.RegistrationDTO{
		Name:     "new_name",
		Email:    "new@mail",
		Password: "test_password",
		ClassNum: "111011",
	})

	resp, err := http.Post(MainURL+"/student/registration", contentType, bytes.NewBuffer(body))
	assert.NilError(t, err)
	defer resp.Body.Close()
	expected, _ := json.Marshal(map[string]string{"error": "unable to found student class: <nil>"})
	recieved, _ := io.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, 400)
	assert.Equal(t, string(expected), string(recieved))
}

func TestRegistrationInvalidEmailFormat(t *testing.T) {
	body, _ := json.Marshal(dto.RegistrationDTO{
		Name:     "new_name",
		Email:    "newmail",
		Password: "test_password",
		ClassNum: "111111",
	})

	resp, err := http.Post(MainURL+"/student/registration", contentType, bytes.NewBuffer(body))
	assert.NilError(t, err)
	defer resp.Body.Close()
	expected, _ := json.Marshal(map[string]string{"error": "unable to create student: pq: value for domain email_type violates check constraint \"email_type_check\""})
	recieved, _ := io.ReadAll(resp.Body)
	assert.Equal(t, resp.StatusCode, 400)
	assert.Equal(t, string(expected), string(recieved))
}

func TestLogout(t *testing.T) {
	body, _ := json.Marshal(nil)
	resp, err := http.Post(MainURL+"/student/logout", contentType, bytes.NewBuffer(body))
	assert.NilError(t, err)
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
