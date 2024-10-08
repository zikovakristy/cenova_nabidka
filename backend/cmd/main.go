package main

import (
    "cenova_nabidka/backend/internal/config"
    "cenova_nabidka/backend/internal/database"
    "cenova_nabidka/backend/internal/export"
    "cenova_nabidka/backend/internal/offer"
    "cenova_nabidka/backend/internal/product"
    "cenova_nabidka/backend/internal/user"
    "cenova_nabidka/backend/pkg/middleware"

    "github.com/gin-gonic/gin"
)

func main() {
    // Inicializace databáze
    config.SetupDatabase()

    // Inicializace repozitářů
    userRepo := user.NewRepository(database.DB)
    productRepo := product.NewRepository(database.DB)
    offerRepo := offer.NewRepository(database.DB)

    // Inicializace služeb
    userService := user.NewService(userRepo)
    productService := product.NewService(productRepo)
    offerService := offer.NewService(offerRepo)

    // Inicializace handlerů
    userHandler := user.NewHandler(userService)
    productHandler := product.NewHandler(productService)
    offerHandler := offer.NewHandler(offerService)

    // Inicializace Gin routeru
    router := gin.Default()

    // Middleware (např. autentizace)
    router.Use(middleware.AuthMiddleware())

    // Registrace rout
    api := router.Group("/api")
    {
        userHandler.RegisterRoutes(api)
        productHandler.RegisterRoutes(api)
        offerHandler.RegisterRoutes(api)
    }

    // Spuštění serveru
    router.Run(":8081")
}
