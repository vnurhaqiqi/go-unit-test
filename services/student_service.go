package services

import (
	"errors"
	"go-unit-test/entities"
	"go-unit-test/repositories"
)

type StudentService struct {
	Repository repositories.StudentRepository
}

func (service StudentService) Get(id int) (*entities.Student, error) {
	student := service.Repository.FindById(id)

	if student == nil {
		return student, errors.New("Student not found")
	} else {
		return student, nil
	}
}