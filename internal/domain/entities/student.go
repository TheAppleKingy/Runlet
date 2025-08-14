package entities

type Student struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string
	ClassId  int      `json:"class_id" db:"class_id"`
	Courses  []Course `json:"courses,omitempty" db:"-"`
}
