package database

import (
	"GoProject/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Veritabanı bağlantısını başlatma
func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Veritabanına bağlanılamadı!")
	}
	DB.AutoMigrate(&models.User{}) // Kullanıcı modeli ile veritabanı şemasını senkronize et
}
