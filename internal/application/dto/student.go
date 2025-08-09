package dto

type StudentViewDTO struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	ClassId int             `json:"class_id"`
	Courses []CourseViewDTO `json:"courses"`
}

// func GetStudentViewDTO(orm *ent.Student) StudentViewDTO {
// 	dto := StudentViewDTO{
// 		Id:    orm.ID,
// 		Name:  orm.Name,
// 		ClassId: orm.ClassID,
// 	}
// 	for _, problem := range orm.Edges.Problems {
// 		dto.Problems = append(dto.Problems, GetProblemView(problem))
// 	}
// 	return dto
// }
