package repositories

import "go-unit-test/entities"

type (
	StudentRepository interface {
		FindById(id int) *entities.Student
	}
)