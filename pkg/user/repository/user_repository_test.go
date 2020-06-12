package repository_test

import (
	"go-mysql-api/pkg/infrastructure/database"
	"go-mysql-api/pkg/user/repository"
	"regexp"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func setUpMockDB() (sqlmock.Sqlmock, *database.DB) {
	// gormのデフォルトテーブル名のルールを修正
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return strings.Replace(defaultTableName, "_data_table", "", 1)
	}
	d, mock, _ := sqlmock.New()
	db := new(database.DB)
	db.Connection, _ = gorm.Open("mysql", d)

	return mock, db
}

func TestFindByID(t *testing.T) {

	mock, db := setUpMockDB()

	query := regexp.QuoteMeta("SELECT * FROM `users` WHERE (`users`.`id` = 1) ORDER BY `users`.`id` ASC LIMIT 1")
	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at"}).
		AddRow(1, "mock user", "mock@mock.com", time.Now(), time.Now())

	mock.ExpectQuery(query).WillReturnRows(rows)

	userRepository := repository.NewUserRepository(db)

	user, err := userRepository.FindByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
