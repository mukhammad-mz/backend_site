package db

import (
	"github.com/jinzhu/gorm"
	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysqlDB *gorm.DB

//ConnectDB connection in DB
func ConnectDB() error {
	var err error
	mysqlDB, err = gorm.Open("mysql", "root:@/sitedb?charset=utf8&parseTime=True&loc=Local")
	return err
}

// GetDB get DB connection
func GetDB() *gorm.DB {
	return mysqlDB
}
