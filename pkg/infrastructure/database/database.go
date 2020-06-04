package database

import (
	"go-mysql-api/pkg/infrastructure/config"
	"strings"

	"github.com/jinzhu/gorm"
)

// DB Database
type DB struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	Connection *gorm.DB
}

// NewDB DataBase create
func NewDB() *DB {
	return newDB(&DB{
		Host:     config.DataBaseConfig.Host,
		Username: config.DataBaseConfig.User,
		Password: config.DataBaseConfig.Password,
		DBName:   config.DataBaseConfig.Database,
	})
}

func newDB(d *DB) *DB {
	db, err := gorm.Open("mysql", d.Username+":"+d.Password+"@tcp("+d.Host+")/"+d.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}

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

// Connect connect a database
func (db *DB) Close() *gorm.DB {
	return db.Connection
}
