package dto

type TeacherViewDTO struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Classes []ClassViewDTO
	Courses []CourseViewDTO
}
