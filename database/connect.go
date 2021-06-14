package database

import (
	"github.com/DucHoangManh/go-admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
func Connect() {
	db ,err := gorm.Open(mysql.Open("mysqluser:1@/go_admin"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	DB = db
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})
}