package repository

import (
	"go-mysql-api/pkg/domain"
	"go-mysql-api/pkg/infrastructure/database"

	"github.com/jinzhu/gorm"
)

// UserRepository repository
type UserRepository interface {
	GetAll() ([]domain.User_DataTable, error)
	GetByID(id int) (domain.User_DataTable, error)
	Create(user domain.User) error
	Update(user domain.User) error
	Delete(id int) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository is init for UserController
func NewUserRepository(db *database.DB) UserRepository {
	return &userRepository{
		db: db.Connection,
	}
}

// GetAll Get all usersdata
func (r *userRepository) GetAll() ([]domain.User_DataTable, error) {
	users := []domain.User_DataTable{}
	err := r.db.Find(&users).Error

	return users, err
}

// GetByID Get single usersdata
func (r *userRepository) GetByID(id int) (domain.User_DataTable, error) {
	user := domain.User_DataTable{}
	err := r.db.First(&user, id).Error

	return user, err
}

// Create Add user
func (r *userRepository) Create(user domain.User) error {
	return r.db.Create(&user).Error
}

// Update Update user
func (r *userRepository) Update(user domain.User) error {
	targetUser := domain.User{}
	if err := r.db.First(&targetUser, user.ID).Error; err != nil {
		return err
	}

	return r.db.Save(&user).Error
}

// Delete Delete userdata
func (r *userRepository) Delete(id int) error {
	user := domain.User_DataTable{}

	if id <= 0 {
		return nil
	}

	user.ID = id

	return r.db.Delete(&user).Error
}
