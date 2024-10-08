package offer

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type Handler struct {
    service Service
}

func NewHandler(service Service) *Handler {
    return &Handler{service}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
    r.POST("/offers", h.CreateOffer)
    r.GET("/offers", h.GetAllOffers)
    r.GET("/offers/:id", h.GetOfferByID)
    r.GET("/users/:user_id/offers", h.GetOffersByUserID)
    r.PUT("/offers/:id", h.UpdateOffer)
    r.DELETE("/offers/:id", h.DeleteOffer)
}

func (h *Handler) CreateOffer(c *gin.Context) {
    var offer PriceOffer
    if err := c.ShouldBindJSON(&offer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.service.CreateOffer(&offer); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, offer)
}

func (h *Handler) GetAllOffers(c *gin.Context) {
    offers, err := h.service.GetAllOffers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, offers)
}

func (h *Handler) GetOfferByID(c *gin.Context) {
    id := c.Param("id")
    offerID, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offer ID"})
        return
    }
    offer, err := h.service.GetOfferByID(uint(offerID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Offer not found"})
        return
    }
    c.JSON(http.StatusOK, offer)
}

func (h *Handler) GetOffersByUserID(c *gin.Context) {
    userID := c.Param("user_id")
    uid, err := strconv.ParseUint(userID, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }
    offers, err := h.service.GetOffersByUserID(uint(uid))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, offers)
}

func (h *Handler) UpdateOffer(c *gin.Context) {
    id := c.Param("id")
    offerID, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offer ID"})
        return
    }

    var offer PriceOffer
    if err := c.ShouldBindJSON(&offer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    offer.ID = uint(offerID)

    if err := h.service.UpdateOffer(&offer); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, offer)
}

func (h *Handler) DeleteOffer(c *gin.Context) {
    id := c.Param("id")
    offerID, err := strconv.ParseUint(id, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offer ID"})
        return
    }

    if err := h.service.DeleteOffer(uint(offerID)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Offer deleted successfully"})
}
