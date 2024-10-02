package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"` // Birincil anahtar
	Name     string `json:"name"`                 // Kullanıcı adı
	Email    string `json:"email"`                // E-posta
	Password string `json:"password"`             // Şifre
}

// Kullanıcıyı veritabanına kaydet
func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

// Tüm kullanıcıları getir
func GetUsers(db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// ID'ye göre kullanıcıyı getir
func GetUserByID(db *gorm.DB, id string) (User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Kullanıcıyı güncelle
func UpdateUser(db *gorm.DB, user *User) error {
	return db.Save(user).Error
}

// Kullanıcıyı sil
func DeleteUser(db *gorm.DB, user *User) error {
	return db.Delete(user).Error
}
