package entities

type Attempt struct {
	StudentId int       `json:"student_id"`
	ProblemId int       `json:"problem_id"`
	Amount    int       `json:"amount"`
	Done      bool      `json:"done"`
	TestCases TestCases `json:"test_cases"`
}
