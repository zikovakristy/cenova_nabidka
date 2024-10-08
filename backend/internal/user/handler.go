package user

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type Handler struct {
    service Service
}

func NewHandler(service Service) *Handler {
    return &Handler{service}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
    r.POST("/users", h.CreateUser)
    r.GET("/users", h.GetAllUsers)
    r.GET("/users/:id", h.GetUserByID)
    r.PUT("/users/:id", h.UpdateUser)
    r.DELETE("/users/:id", h.DeleteUser)
}

func (h *Handler) CreateUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.service.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, user)
}

func (h *Handler) GetAllUsers(c *gin.Context) {
    users, err := h.service.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUserByID(c *gin.Context) {
    id := c.Param("id")
    var userID uint
    if _, err := fmt.Sscanf(id, "%d", &userID); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
    user, err := h.service.GetUserByID(userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var userID uint
    if _, err := fmt.Sscanf(id, "%d", &userID); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user.ID = userID

    if err := h.service.UpdateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(c *gin.Context) {
    id := c.Param("id")
    var userID uint
    if _, err := fmt.Sscanf(id, "%d", &userID); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    if err := h.service.DeleteUser(userID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
