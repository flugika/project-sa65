package controller

import (
	"net/http"

	"github.com/flugika/project-sa65/entity"
	"github.com/gin-gonic/gin"
)

// POST /students
func CreateStudent(c *gin.Context) {
	var student entity.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": student})
}

// GET /student/:id
func GetStudent(c *gin.Context) {
	var student entity.Student
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM students WHERE id = ?", id).Find(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": student})
}

// GET /students
func ListStudents(c *gin.Context) {
	var students []entity.Student
	if err := entity.DB().Raw("SELECT * FROM students").Find(&students).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": students})
}

// DELETE /students/:id
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM students WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "student not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /students
func UpdateStudent(c *gin.Context) {
	var student entity.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", student.ID).First(&student); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "student not found"})
		return
	}

	if err := entity.DB().Save(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": student})
}