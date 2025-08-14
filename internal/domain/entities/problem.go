package entities

type Problem struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CourseId    int       `json:"course_id"`
	Attempts    []Attempt `json:"attempts,omitempty" db:"-"`
	Students    []Student `json:"students,omitempty" db:"-"`
}
