package repository_test

import (
	"go-mysql-api/pkg/domain"
	"go-mysql-api/pkg/infrastructure/database"
	"go-mysql-api/pkg/user/repository"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func setUpMockDB() (sqlmock.Sqlmock, *database.DB) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return strings.Replace(defaultTableName, "_data_table", "", 1)
	}
	d, mock, _ := sqlmock.New()
	db := new(database.DB)
	// gorm+sqlmockの場合テストが読みづらくなるのでmockのdriverははpostgresにする
	db.Connection, _ = gorm.Open("mysql", d)

	return mock, db
}

func TestFindAll(t *testing.T) {
	mock, db := setUpMockDB()

	query := regexp.QuoteMeta("SELECT * FROM `users`")
	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"}).
		AddRow(1, "mock user", "mock@mock.com", time.Now(), time.Now())
	mock.ExpectQuery(query).WillReturnRows(rows)

	userRepository := repository.NewUserRepository(db)

	user, err := userRepository.FindAll()
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestFindByID(t *testing.T) {
	mock, db := setUpMockDB()

	id := 1

	query := regexp.QuoteMeta("SELECT * FROM `users` WHERE (`users`.`id` = " + strconv.Itoa(id) +
		") ORDER BY `users`.`id` ASC LIMIT 1")
	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"}).
		AddRow(1, "mock user", "mock@mock.com", time.Now(), time.Now())
	mock.ExpectQuery(query).WillReturnRows(rows)

	userRepository := repository.NewUserRepository(db)

	user, err := userRepository.FindByID(id)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestCreate(t *testing.T) {
	mock, db := setUpMockDB()

	mockUser := domain.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"

	mock.ExpectBegin()
	query := regexp.QuoteMeta("INSERT INTO `users` (`id`,`name`,`email`) VALUES (?,?,?)")
	mock.ExpectExec(query).
		WithArgs(mockUser.ID, mockUser.Name, mockUser.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	userRepository := repository.NewUserRepository(db)

	err := userRepository.Create(mockUser)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	mock, db := setUpMockDB()

	mockUser := domain.User{}
	mockUser.ID = 1
	mockUser.Name = "mockuser"
	mockUser.Email = "mock@mock.com"

	selectQuery := regexp.QuoteMeta("SELECT * FROM `users` WHERE (`users`.`id` = " +
		strconv.Itoa(mockUser.ID) +
		") ORDER BY `users`.`id` ASC LIMIT 1")
	selectRows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"}).
		AddRow(1, "mock user", "mock@mock.com", time.Now(), time.Now())
	mock.ExpectQuery(selectQuery).WillReturnRows(selectRows)

	mock.ExpectBegin()
	query := regexp.QuoteMeta("UPDATE `users` SET `name` = ?, `email` = ? WHERE `users`.`id` = ?")
	mock.ExpectExec(query).
		WithArgs(mockUser.Name, mockUser.Email, mockUser.ID).
		WillReturnResult(sqlmock.NewResult(int64(mockUser.ID), 1))
	mock.ExpectCommit()

	userRepository := repository.NewUserRepository(db)

	err := userRepository.Update(mockUser)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	mock, db := setUpMockDB()
	id := 1

	mock.ExpectBegin()
	query := regexp.QuoteMeta("DELETE FROM `users` WHERE `users`.`id` = ?")
	mock.ExpectExec(query).WithArgs(id).WillReturnResult(sqlmock.NewResult(int64(id), 1))
	mock.ExpectCommit()

	userRepository := repository.NewUserRepository(db)

	err := userRepository.Delete(1)
	assert.NoError(t, err)
}
