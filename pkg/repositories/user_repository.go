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
	db.Find(&users)

	return users, db.Error
}

// FindByID Get single usersdata
func (r *UserRepository) FindByID(id int) (result models.User_DataTable, err error) {
	d := database.NewDB()
	db := d.Connect()
	defer d.Close()

	user := models.User_DataTable{}
	db.First(&user, id)

	return user, db.Error
}

// Regist Add user
func (r *UserRepository) Regist(user models.User) error {
	d := database.NewDB()
	db := d.Connect()
	defer d.Close()

	db.Create(&user)

	return db.Error
}

// Update Update user
func (r *UserRepository) Update(user models.User) error {
	d := database.NewDB()
	db := d.Connect()
	defer d.Close()

	db.Save(&user)

	return db.Error
}

// Delete Delete userdata
func (r *UserRepository) Delete(id int) error {
	d := database.NewDB()
	db := d.Connect()
	defer d.Close()

	user := models.User_DataTable{}
	user.ID = id

	db.Delete(&user)

	return db.Error
}
