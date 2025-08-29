package dto

type ProblemForCourse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type CodeSolution struct {
	Lang string `json:"lang" binding:"required"`
	Code string `json:"code" binding:"required"`
}
