package dto

type ClassViewDTO struct {
	Id       int              `json:"id"`
	Number   string           `json:"number"`
	Students []StudentViewDTO `json:"students"`
	Teachers []TeacherViewDTO `json:"teachers"`
	Courses  []CourseViewDTO  `json:"courses"`
}
