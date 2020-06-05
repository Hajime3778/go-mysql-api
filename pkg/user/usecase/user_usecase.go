package usecase

import (
	"go-mysql-api/pkg/domain"
	"go-mysql-api/pkg/user/repository"
)

// UserUsecase usecase
type UserUsecase interface {
	GetUsers() ([]domain.User_DataTable, error)
	GetUser(id int) (domain.User_DataTable, error)
	CreateUser(user domain.User) error
	UpdateUser(user domain.User) error
	DeleteUser(id int) error
}

// userUsecase usecase
type userUsecase struct {
	repo repository.UserRepository
}

// NewUserUsecase is init for UserUsecase
func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

// GetUsers 複数のUserを取得します
func (u *userUsecase) GetUsers() ([]domain.User_DataTable, error) {
	return u.repo.GetAll()
}

// GetUser 1件のUserを取得します
func (u *userUsecase) GetUser(id int) (domain.User_DataTable, error) {
	return u.repo.FindByID(id)
}

// CreateUser Userを作成します
func (u *userUsecase) CreateUser(user domain.User) error {
	return u.repo.Regist(user)
}

// UpdateUser Userを更新します。
func (u *userUsecase) UpdateUser(user domain.User) error {
	return u.repo.Update(user)
}

// DeleteUser Userを削除します
func (u *userUsecase) DeleteUser(id int) error {
	return u.repo.Delete(id)
}
