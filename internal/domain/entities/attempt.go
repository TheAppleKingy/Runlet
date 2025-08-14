package entities

type Attempt struct {
	ID        int  `json:"id"`
	Amount    uint `json:"amount"`
	Done      bool `json:"done"`
	StudentId int  `json:"student_id"`
	ProblemId int  `json:"problem_id"`
}
