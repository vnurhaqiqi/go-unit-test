package services

import (
	"go-unit-test/entities"
	"go-unit-test/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var studentRepository = &repositories.StudentRepositoryMock{Mock: mock.Mock{}}
var studentService = StudentService{Repository: studentRepository}

func TestStudentService_Get(t *testing.T) {
	// do mock

	studentRepository.Mock.On("FindById", 1).Return(nil)

	student, err := studentService.Get(1)
	assert.Nil(t, student)
	assert.NotNil(t, err) 
}

func TestStudentService_GetSuccess(t *testing.T) {
	student := entities.Student{
		ID: 1,
		Name: "Viqi Nurhaqiqi",
	}

	studentRepository.Mock.On("FindById", 2).Return(nil)

	result, err := studentService.Get(2)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, student.ID, result.ID)
	assert.Equal(t, student.Name, result.Name)
}