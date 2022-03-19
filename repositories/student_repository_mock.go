package repositories

import (
	"go-unit-test/entities"

	"github.com/stretchr/testify/mock"
)

type StudentRepositoryMock struct {
	Mock mock.Mock
}

func (repository *StudentRepositoryMock) FindById(id int) *entities.Student {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	} else {
		student := arguments.Get(0).(entities.Student)

		return &student
	}
}