package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/api/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Ahoj, svět!",
        })
    })

    router.Run(":8081")
}
