package dto

import "Runlet/internal/infrastructure/ent"

type AttemptViewDTO struct {
	Id        int  `json:"id"`
	Amount    uint `json:"amount"`
	Done      bool `json:"done"`
	StudentId int  `json:"student_id"`
	ProblemId int  `json:"problem_id"`
}

func GetAttemptViewDTO(orm *ent.Attempt) AttemptViewDTO {
	return AttemptViewDTO{
		Amount:    orm.Amount,
		Done:      orm.Done,
		StudentId: orm.StudentID,
		ProblemId: orm.ProblemID,
	}
}
