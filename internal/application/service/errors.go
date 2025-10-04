package service

import "errors"

var (
	ErrStudentNotFound    = errors.New("unable to find student")
	ErrRegisterStudent    = errors.New("unable to register student")
	ErrStudentHasNoCourse = errors.New("provided student does not belong to provided course")
)

var (
	ErrWrongPassword      = errors.New("wrong password")
	ErrCreateToken        = errors.New("unable to create token")
	ErrProcessingPassword = errors.New("error processing password")
)

var (
	ErrTeacherNotFound = errors.New("unable to find teacher")
	ErrRegisterTeacher = errors.New(("unable to register teacher"))
)

var (
	ErrClassNotFound = errors.New("unable to find class")
)
