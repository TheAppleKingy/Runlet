package dto

type ProblemView struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CourseId    int    `json:"course_id" binding:"required"`
}
