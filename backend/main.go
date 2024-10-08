package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// Definice modelu
type Product struct {
    ID    uint    `gorm:"primaryKey" json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

func main() {
    // Načtení environment variables z .env souboru
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Konfigurace databáze
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbSSLMode := os.Getenv("DB_SSLMODE")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Europe/Prague",
        dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Migrace databázových modelů
    db.AutoMigrate(&Product{})

    router := gin.Default()

    // Endpoint pro získání všech produktů
    router.GET("/api/products", func(c *gin.Context) {
        var products []Product
        db.Find(&products)
        c.JSON(http.StatusOK, products)
    })

    // Endpoint pro vytvoření nového produktu
    router.POST("/api/products", func(c *gin.Context) {
        var product Product
        if err := c.ShouldBindJSON(&product); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        db.Create(&product)
        c.JSON(http.StatusCreated, product)
    })

    // Endpoint pro získání jednoho produktu podle ID
    router.GET("/api/products/:id", func(c *gin.Context) {
        var product Product
        id := c.Param("id")
        if err := db.First(&product, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Produkt nenalezen"})
            return
        }
        c.JSON(http.StatusOK, product)
    })

    // Endpoint pro aktualizaci produktu podle ID
    router.PUT("/api/products/:id", func(c *gin.Context) {
        var product Product
        id := c.Param("id")
        if err := db.First(&product, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Produkt nenalezen"})
            return
        }
        if err := c.ShouldBindJSON(&product); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        db.Save(&product)
        c.JSON(http.StatusOK, product)
    })

    // Endpoint pro smazání produktu podle ID
    router.DELETE("/api/products/:id", func(c *gin.Context) {
        var product Product
        id := c.Param("id")
        if err := db.First(&product, id).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Produkt nenalezen"})
            return
        }
        db.Delete(&product)
        c.JSON(http.StatusOK, gin.H{"message": "Produkt smazán"})
    })

    // Spuštění serveru na portu 8081
    router.Run(":8081")
}
