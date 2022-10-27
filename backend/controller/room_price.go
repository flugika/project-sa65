package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/flugika/project-sa65/entity"
)

// POST /room_prices
func CreateRoom_Price(c *gin.Context) {
	var room_price entity.Room_Price
	if err := c.ShouldBindJSON(&room_price); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&room_price).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": room_price})
}

// GET /room_price/:id
func GetRoom_Price(c *gin.Context) {
	var room_price entity.Room_Price
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM room_prices WHERE id = ?", id).Find(&room_price).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if room_price.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": room_price})
}

// GET /room_prices
func ListRoom_Prices(c *gin.Context) {
	var room_prices []entity.Room_Price
	if err := entity.DB().Raw("SELECT * FROM room_prices").Find(&room_prices).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": room_prices})
}

// DELETE /room_prices/:id
func DeleteRoom_Price(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM room_prices WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room_price not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /room_prices
func UpdateRoom_Price(c *gin.Context) {
	var room_price entity.Room_Price
	if err := c.ShouldBindJSON(&room_price); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", room_price.ID).First(&room_price); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room_price not found"})
		return
	}

	if err := entity.DB().Save(&room_price).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": room_price})
}