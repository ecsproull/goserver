package handlers

import (
	"goserver/internal/models"
	"goserver/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlaceHandler struct{}

func NewPlaceHandler() *PlaceHandler {
	return &PlaceHandler{}
}

// GET /api/v1/places
func (h *PlaceHandler) GetPlaces(c *gin.Context) {
	places, err := services.GetPlaces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, places)
}

// POST /api/v1/places
func (h *PlaceHandler) SavePlace(c *gin.Context) {
	var place models.DbPlace
	if err := c.ShouldBindJSON(&place); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.SavePlace(&place); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Place created successfully"})
}

// PUT /api/v1/places/:id
func (h *PlaceHandler) UpdatePlace(c *gin.Context) {
	id := c.Param("id")

	var place models.DbPlace
	if err := c.ShouldBindJSON(&place); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	placeID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	place.ID = placeID
	updated, err := services.UpdatePlace(&place)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Place updated successfully", "place": updated})
}

// DELETE /api/v1/places/:id
func (h *PlaceHandler) DeletePlace(c *gin.Context) {
	id := c.Param("id")
	deleted, err := services.DeletePlace(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Place deleted successfully", "place": deleted})
}
