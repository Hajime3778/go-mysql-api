package database

import (
	"fmt"
	"go-mysql-api/pkg/infrastructure/config"
	"net/url"
	"strings"

	"github.com/jinzhu/gorm"
)

// DB Database
type DB struct {
	Host       string
	Port       string
	Username   string
	Password   string
	DBName     string
	Connection *gorm.DB
}

// NewDB DataBase create
func NewDB(c *config.Config) *DB {
	return newDB(&DB{
		Host:     c.DataBase.Host,
		Port:     c.DataBase.Port,
		Username: c.DataBase.User,
		Password: c.DataBase.Password,
		DBName:   c.DataBase.Database,
	})
}

func newDB(d *DB) *DB {
	connectionInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.DBName)

	option := url.Values{}
	option.Add("charset", "utf8")
	option.Add("parseTime", "True")
	option.Add("loc", "Local")

	connection := fmt.Sprintf("%s?%s", connectionInfo, option.Encode())

	db, err := gorm.Open("mysql", connection)
	if err != nil {
		panic(err.Error())
	}

	// gormのデフォルトテーブル名のルールを修正
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return strings.Replace(defaultTableName, "_data_table", "", 1)
	}

	d.Connection = db
	return d
}

// Begin begins a transaction
func (db *DB) Begin() *gorm.DB {
	return db.Connection.Begin()
}

// Connect connect a database
func (db *DB) Connect() *gorm.DB {
	return db.Connection
}

// Close close a database
func (db *DB) Close() *gorm.DB {
	return db.Connection
}
