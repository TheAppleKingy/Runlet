package entities

type Problem struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CourseId    int       `json:"course_id,omitempty" db:"course_id"`
	TestCases   TestCases `json:"test_cases,omitempty" db:"test_cases"`
	Attempts    []Attempt `json:"attempts,omitempty" db:"-"`
	Students    []Student `json:"students,omitempty" db:"-"`
}
