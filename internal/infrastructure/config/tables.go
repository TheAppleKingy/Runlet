package config

type tables struct {
	Student         string
	Teacher         string
	Course          string
	Problem         string
	Class           string
	Attempt         string
	ClassesCourses  string
	TeachersClasses string
}

// Tables store names of database tables for pretty using in query builders
var Tables = tables{
	Student:         "students",
	Teacher:         "teachers",
	Course:          "courses",
	Problem:         "problems",
	Class:           "classes",
	Attempt:         "attempts",
	ClassesCourses:  "classes_courses",
	TeachersClasses: "teachers_classes",
}
