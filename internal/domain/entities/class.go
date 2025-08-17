package entities

type Class struct {
	ID       int       `json:"id"`
	Number   string    `json:"number"`
	Students []Student `json:"students,omitempty" db:"-"`
	Teachers []Teacher `json:"teachers,omitempty" db:"-"`
	Courses  []Course  `json:"courses,omitempty" db:"-"`
}
