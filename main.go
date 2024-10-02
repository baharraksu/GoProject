package main

import (
	"GoProject/database" // Veritabanı işlemleri için import
	"GoProject/routes"   // Rota ayarları için import

	"github.com/gin-gonic/gin" // Gin framework'ü için import
)

func main() {
	r := gin.Default() // Gin router'ı oluştur

	// Veritabanını başlat
	database.InitDatabase()

	// Rotaları ayarla
	routes.SetupRoutes(r, database.DB)

	// Sunucuyu başlat (8080 portu üzerinde)
	if err := r.Run(":8080"); err != nil {
		panic("Sunucu başlatılamadı: " + err.Error())
	}
}
