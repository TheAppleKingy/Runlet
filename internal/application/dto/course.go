package dto

import (
	"Runlet/internal/infrastructure/ent"
)

type CourseViewDTO struct {
	Id          int              `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Problems    []ProblemViewDTO `json:"problems,omitempty"`
	Classes     []ClassViewDTO   `json:"classes,omitempty"`
	Teachers    []TeacherViewDTO `json:"teachers,omitempty"`
}

func GetCourseViewNoEdgesDTO(orm *ent.Course) CourseViewDTO {
	return CourseViewDTO{
		Id:          orm.ID,
		Title:       orm.Title,
		Description: orm.Description,
	}
}

func GetCourseForStudentViewDTO(orm *ent.Course) CourseViewDTO {
	dto := GetCourseViewNoEdgesDTO(orm)
	for _, problem := range orm.Edges.Problems {
		dto.Problems = append(dto.Problems, GetProblemViewNoEdgesDTO(problem))
	}
	return dto
}

func GetCoursesForStudentViewDTO(orms []*ent.Course) []CourseViewDTO {
	var dtos []CourseViewDTO
	for _, orm := range orms {
		dtos = append(dtos, GetCourseForStudentViewDTO(orm))
	}
	return dtos
}

type CourseCreateDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ClassesIds  []int  `json:"classes_ids" binding:"required"`
	TeachersIds []int  `json:"teachers_ids" binding:"required"`
}
