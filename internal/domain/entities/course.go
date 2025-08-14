package entities

type Course struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Problems    []Problem `json:"problems,omitempty" db:"-"`
	Classes     []Class   `json:"classes,omitempty" db:"-"`
	Teachers    []Teacher `json:"teachers,omitempty" db:"-"`
}
