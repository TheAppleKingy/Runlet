package repositoryimpl

import (
	"Runlet/internal/infrastructure/ent"
)

func NewCourseRepository(dbClient *ent.Client) *CourseRepository {
	return &CourseRepository{
		client: dbClient,
	}
}

func NewStudentRepository(dbClient *ent.Client) *StudentRepository {
	return &StudentRepository{
		client: dbClient,
	}
}

func NewClassRepository(dbclient *ent.Client) *ClassRepository {
	return &ClassRepository{
		client: dbclient,
	}
}

func NewTeacherRepository(dbClient *ent.Client) *TeacherRepository {
	return &TeacherRepository{
		client: dbClient,
	}
}

func NewProblemRepository(dbClient *ent.Client) *ProblemRepository {
	return &ProblemRepository{
		client: dbClient,
	}
}
