package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Definování endpointu /api/hello
    router.GET("/api/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Ahoj z backendu!",
        })
    })

    // Spuštění serveru na portu 8081
    router.Run(":8081")
}
