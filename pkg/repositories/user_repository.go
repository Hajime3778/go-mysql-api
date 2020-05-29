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
func (r *UserRepository) GetAll() (result []models.User_DataTable, err error) {
	d := database.NewDB()
	db := d.Connect()
	defer d.Close()

	users := []models.User_DataTable{}
	dbError := db.Find(&users).Error

	return users, dbError
}

// FindByID Get single usersdata
func (r *UserRepository) FindByID(id int) (result models.User_DataTable, err error) {
	d := database.NewDB()
	db := d.Connect()
	defer d.Close()

	user := models.User_DataTable{}
	dbError := db.First(&user, id).Error

	return user, dbError
}

// Regist Add user
func (r *UserRepository) Regist(user models.User) error {
	d := database.NewDB()
	db := d.Connect()
	defer d.Close()

	dbError := db.Create(&user).Error

	return dbError
}

// Update Update user
func (r *UserRepository) Update(user models.User) error {
	d := database.NewDB()
	db := d.Connect()
	defer d.Close()

	dbError := db.Save(&user).Error

	return dbError
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
	dbError := db.Delete(&user).Error

	return dbError
}
