package controllers

import (
	"GoProject/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Yeni kullanıcı oluşturma
func CreateUser(c *gin.Context, db *gorm.DB) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.CreateUser(db, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Tüm kullanıcıları listeleme
func GetUsers(c *gin.Context, db *gorm.DB) {
	users, err := models.GetUsers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Belirli kullanıcıyı getirme
func GetUserByID(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	user, err := models.GetUserByID(db, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kullanıcı bulunamadı"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context, db *gorm.DB) {
	id := c.Param("id") // ID'yi URL parametrelerinden al
	var user models.User

	// ID'ye göre kullanıcıyı al
	existingUser, err := models.GetUserByID(db, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kullanıcı bulunamadı"})
		return
	}

	// Gelen JSON verisini kullanıcı modeline bağla
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Güncelleme işlemi
	existingUser.Name = user.Name   // Var olan kullanıcının adını güncelle
	existingUser.Email = user.Email // Var olan kullanıcının e-posta adresini güncelle

	if err := models.UpdateUser(db, &existingUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Güncellenen kullanıcıyı yanıt olarak döndür
	c.JSON(http.StatusOK, existingUser)
}

// Kullanıcıyı silme
func DeleteUser(c *gin.Context, db *gorm.DB) {
	id := c.Param("id") // ID'yi URL parametrelerinden al

	// ID'ye göre kullanıcıyı al
	existingUser, err := models.GetUserByID(db, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kullanıcı bulunamadı"})
		return
	}

	// Kullanıcıyı sil
	if err := models.DeleteUser(db, &existingUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Başarı mesajı döndür
	c.JSON(http.StatusOK, gin.H{"message": "Kullanıcı başarıyla silindi"})
}
