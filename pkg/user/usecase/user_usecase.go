package usecase

import (
	"go-mysql-api/pkg/domain"
	"go-mysql-api/pkg/user/repository"
)

// UserUsecase usecase
type UserUsecase interface {
	FindAll() ([]domain.User_DataTable, error)
	FindByID(id int) (domain.User_DataTable, error)
	Create(user domain.User) error
	Update(user domain.User) error
	Delete(id int) error
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

// FindAll 複数のUserを取得します
func (u *userUsecase) FindAll() ([]domain.User_DataTable, error) {
	return u.repo.FindAll()
}

// FindByID 1件のUserを取得します
func (u *userUsecase) FindByID(id int) (domain.User_DataTable, error) {
	return u.repo.FindByID(id)
}

// Create Userを作成します
func (u *userUsecase) Create(user domain.User) error {
	return u.repo.Create(user)
}

// Update Userを更新します。
func (u *userUsecase) Update(user domain.User) error {
	return u.repo.Update(user)
}

// Delete Userを削除します
func (u *userUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}
