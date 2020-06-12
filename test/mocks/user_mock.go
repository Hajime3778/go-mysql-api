package mocks

import (
	"go-mysql-api/pkg/domain"

	"github.com/stretchr/testify/mock"
)

// UserUsecase is mock
type UserUsecase struct {
	mock.Mock
}

// FindAll is mock function
func (_m *UserUsecase) FindAll() ([]domain.User_DataTable, error) {
	ret := _m.Called()
	return ret.Get(0).([]domain.User_DataTable), ret.Error(1)
}

// FindByID is mock function
func (_m *UserUsecase) FindByID(id int) (domain.User_DataTable, error) {
	ret := _m.Called(id)
	return ret.Get(0).(domain.User_DataTable), ret.Error(1)
}

// Create is mock function
func (_m *UserUsecase) Create(user domain.User) error {
	ret := _m.Called(user)
	return ret.Error(0)
}

// Update is mock function
func (_m *UserUsecase) Update(user domain.User) error {
	ret := _m.Called(user)
	return ret.Error(0)
}

// Delete is mock function
func (_m *UserUsecase) Delete(id int) error {
	ret := _m.Called(id)
	return ret.Error(0)
}

// UserRepository is mock
type UserRepository struct {
	mock.Mock
}

// FindAll is mock function
func (_m *UserRepository) FindAll() ([]domain.User_DataTable, error) {
	ret := _m.Called()
	return ret.Get(0).([]domain.User_DataTable), ret.Error(1)
}

// FindByID is mock function
func (_m *UserRepository) FindByID(id int) (domain.User_DataTable, error) {
	ret := _m.Called(id)
	return ret.Get(0).(domain.User_DataTable), ret.Error(1)
}

// Create is mock function
func (_m *UserRepository) Create(user domain.User) error {
	ret := _m.Called(user)
	return ret.Error(0)
}

// Update is mock function
func (_m *UserRepository) Update(user domain.User) error {
	ret := _m.Called(user)
	return ret.Error(0)
}

// Delete is mock function
func (_m *UserRepository) Delete(id int) error {
	ret := _m.Called(id)
	return ret.Error(0)
}
