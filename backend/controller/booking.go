package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/flugika/project-sa65/entity"
)

// POST /bookings
func CreateBooking(c *gin.Context) {
	var booking entity.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": booking})
}

// GET /booking/:id
func GetBooking(c *gin.Context) {
	var booking entity.Booking
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&booking); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booking not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": booking})
}

// GET /bookings
func ListBookings(c *gin.Context) {
	var bookings []entity.Booking
	if err := entity.DB().Raw("SELECT * FROM bookings").Scan(&bookings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookings})
}

// DELETE /bookings/:id
func DeleteBooking(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM bookings WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booking not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /bookings
func UpdateBooking(c *gin.Context) {
	var booking entity.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", booking.ID).First(&booking); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booking not found"})
		return
	}

	if err := entity.DB().Save(&booking).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": booking})
}