package repositories

import (
	"go-mysql-api/pkg/database"
	"go-mysql-api/pkg/models"
)

// UserRepository repository
type UserRepository struct{}

// NewUserRepository is init for UserController
func NewUserRepository() *UserRepository {
	return new(UserRepository)
}

// GetAll Get all usersdata
func (r *UserRepository) GetAll() ([]models.User_DataTable, error) {
	d := database.NewDB()
	db := d.Connect()
	defer d.Close()

	users := []models.User_DataTable{}
	err := db.Find(&users).Error

	return users, err
}

// FindByID Get single usersdata
func (r *UserRepository) FindByID(id int) (models.User_DataTable, error) {
	d := database.NewDB()
	db := d.Connect()
	defer d.Close()

	user := models.User_DataTable{}
	err := db.First(&user, id).Error

	return user, err
}

// Regist Add user
func (r *UserRepository) Regist(user models.User) error {
	d := database.NewDB()
	db := d.Connect()
	defer d.Close()

	return db.Create(&user).Error
}

// Update Update user
func (r *UserRepository) Update(user models.User) error {
	d := database.NewDB()
	db := d.Connect()
	defer d.Close()

	if err := db.First(&user, user.ID).Error; err != nil {
		return err
	}

	return db.Save(&user).Error
}

// Delete Delete userdata
func (r *UserRepository) Delete(id int) error {
	d := database.NewDB()
	db := d.Connect()
	defer d.Close()

	user := models.User_DataTable{}

	if id <= 0 {
		return nil
	}

	user.ID = id

	return db.Delete(&user).Error
}
