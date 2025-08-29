package dto

type Login struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	IsStudent bool   `json:"is_student"`
}

type StudentRegistration struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	ClassNum string `json:"class" binding:"required"`
}

type TeacherRegistration struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
