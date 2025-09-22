package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// ConnectDB connects to MySQL using GORM v2
func ConnectDB() error {
	var err error
	dsn := "root:0000@/sitedb?charset=utf8mb4&parseTime=true&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Ping to verify connection
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	if err = sqlDB.Ping(); err != nil {
		return err
	}

	// Optional: Connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return nil
}

// GetDB returns the global DB connection
func GetDB() *gorm.DB {
	return db
}