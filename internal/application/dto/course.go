package dto

type CourseCreateDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ClassesIds  []int  `json:"classes_ids"`
	TeachersIds []int  `json:"teachers_ids"`
}
