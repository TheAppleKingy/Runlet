package entities

type Teacher struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string
	Classes  []Class  `json:"classes,omitempty" db:"-"`
	Courses  []Course `json:"courses,omitempty" db:"-"`
}
