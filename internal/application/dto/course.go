package dto

type CourseCreate struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ClassesIds  []int  `json:"classes_ids"`
	TeachersIds []int  `json:"teachers_ids"`
}

type CourseForStudent struct {
	ID          int                `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Problems    []ProblemForCourse `json:"problems,omitempty"`
}
