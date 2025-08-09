package dto

import "Runlet/internal/infrastructure/ent"

type ProblemViewDTO struct {
	Id          int              `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	CourseId    int              `json:"course_id"`
	Attempts    []AttemptViewDTO `json:"attempts"`
	Students    []StudentViewDTO `json:"students"`
}

func GetProblemViewNoEdgesDTO(orm *ent.Problem) ProblemViewDTO {
	return ProblemViewDTO{
		Id:          orm.ID,
		Title:       orm.Title,
		Description: orm.Description,
		CourseId:    orm.CourseID,
	}
}

type ProblemForStudentViewDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func GetProblemForStudentViewDTO(orm *ent.Problem) ProblemForStudentViewDTO {
	return ProblemForStudentViewDTO{
		Title:       orm.Title,
		Description: orm.Description,
	}
}

func GetProblemForStudentViewsDTO(orms []*ent.Problem) []ProblemForStudentViewDTO {
	var dtos []ProblemForStudentViewDTO
	for _, orm := range orms {
		dtos = append(dtos, GetProblemForStudentViewDTO(orm))
	}
	return dtos
}
