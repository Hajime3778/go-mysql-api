package usecase_test

import (
	"go-mysql-api/pkg/domain"
	"go-mysql-api/pkg/user/usecase"
	"go-mysql-api/test/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFindAll(t *testing.T) {
	mockUsers := make([]domain.User_DataTable, 0)
	mockUser := domain.User_DataTable{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"
	mockUser.CreatedAt = time.Now()
	mockUser.UpdatedAt = time.Now()
	mockUsers = append(mockUsers, mockUser)

	// モック
	mockUserRepo := new(mocks.UserRepository)

	t.Run("test1", func(t *testing.T) {
		// モックの戻り値を設定
		mockUserRepo.On("FindAll").Return(mockUsers, nil).Once()
		// テスト対象(モックを注入)
		usecase := usecase.NewUserUsecase(mockUserRepo)

		users, err := usecase.FindAll()

		assert.NoError(t, err)
		assert.NotNil(t, users)

		mockUserRepo.AssertExpectations(t)
	})
}

func TestFindByID(t *testing.T) {
	mockUser := domain.User_DataTable{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"
	mockUser.CreatedAt = time.Now()
	mockUser.UpdatedAt = time.Now()

	// モック
	mockUserRepo := new(mocks.UserRepository)

	t.Run("test1", func(t *testing.T) {
		// モックの戻り値を設定
		mockUserRepo.On("FindByID", mockUser.ID).Return(mockUser, nil).Once()
		// テスト対象(モックを注入)
		usecase := usecase.NewUserUsecase(mockUserRepo)

		user, err := usecase.FindByID(mockUser.ID)

		assert.NoError(t, err)
		assert.NotNil(t, user)

		mockUserRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mockUser := domain.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"

	// モック
	mockUserRepo := new(mocks.UserRepository)

	t.Run("test1", func(t *testing.T) {
		// モックの戻り値を設定
		mockUserRepo.On("Create", mockUser).Return(nil).Once()
		// テスト対象(モックを注入)
		usecase := usecase.NewUserUsecase(mockUserRepo)

		err := usecase.Create(mockUser)

		assert.NoError(t, err)

		mockUserRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockUser := domain.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"

	// モック
	mockUserRepo := new(mocks.UserRepository)

	t.Run("test1", func(t *testing.T) {
		// モックの戻り値を設定
		mockUserRepo.On("Update", mockUser).Return(nil).Once()
		// テスト対象(モックを注入)
		usecase := usecase.NewUserUsecase(mockUserRepo)

		err := usecase.Update(mockUser)

		assert.NoError(t, err)

		mockUserRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockUser := domain.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"
	// モック
	mockUserRepo := new(mocks.UserRepository)

	t.Run("test1", func(t *testing.T) {
		// モックの戻り値を設定
		mockUserRepo.On("Delete", mockUser.ID).Return(nil).Once()
		// テスト対象(モックを注入)
		usecase := usecase.NewUserUsecase(mockUserRepo)

		err := usecase.Delete(mockUser.ID)

		assert.NoError(t, err)

		mockUserRepo.AssertExpectations(t)
	})
}
