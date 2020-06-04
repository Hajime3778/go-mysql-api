package repositories

import (
	"go-mysql-api/pkg/domain"

	"github.com/jinzhu/gorm"
)

// UserRepository repository
type UserRepository interface {
	GetAll() ([]domain.User_DataTable, error)
	FindByID(id int) (domain.User_DataTable, error)
	Regist(user domain.User) error
	Update(user domain.User) error
	Delete(id int) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository is init for UserController
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// GetAll Get all usersdata
func (r *userRepository) GetAll() ([]domain.User_DataTable, error) {
	// d := database.NewDB()
	// db := d.Connect()
	// defer d.Close()

	users := []domain.User_DataTable{}
	err := r.db.Find(&users).Error

	return users, err
}

// FindByID Get single usersdata
func (r *userRepository) FindByID(id int) (domain.User_DataTable, error) {
	// d := database.NewDB()
	// db := d.Connect()
	// defer d.Close()

	user := domain.User_DataTable{}
	err := r.db.First(&user, id).Error

	return user, err
}

// Regist Add user
func (r *userRepository) Regist(user domain.User) error {
	// d := database.NewDB()
	// db := d.Connect()
	// defer d.Close()

	return r.db.Create(&user).Error
}

// Update Update user
func (r *userRepository) Update(user domain.User) error {
	// d := database.NewDB()
	// db := d.Connect()
	// defer d.Close()

	targetUser := domain.User{}
	if err := r.db.First(&targetUser, user.ID).Error; err != nil {
		return err
	}

	return r.db.Save(&user).Error
}

// Delete Delete userdata
func (r *userRepository) Delete(id int) error {
	// d := database.NewDB()
	// db := d.Connect()
	// defer d.Close()

	user := domain.User_DataTable{}

	if id <= 0 {
		return nil
	}

	user.ID = id

	return r.db.Delete(&user).Error
}
