package controller

import (
	"net/http"

	"github.com/flugika/project-sa65/entity"
	"github.com/gin-gonic/gin"
)

// POST /payment_bills
func CreatePayment_Bill(c *gin.Context) {

	var payment_bill entity.Payment_Bill
	var employee entity.Employee
	var booking entity.Booking
	var semester entity.Semester

	if err := c.ShouldBindJSON(&payment_bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 10: ค้นหา employee ด้วย id
	if tx := entity.DB().Where("id = ?", payment_bill.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 11: ค้นหา booking ด้วย id
	if tx := entity.DB().Where("id = ?", payment_bill.BookingID).First(&booking); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "booking not found"})
		return
	}

	// 12: ค้นหา semester ด้วย id
	if tx := entity.DB().Where("id = ?", payment_bill.SemesterID).First(&semester); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "semester not found"})
		return
	}

	// 13: สร้าง Payment_Bill
	pb := entity.Payment_Bill{
		Billing_Date:    payment_bill.Billing_Date,    // ตั้งค่าฟิลด์ Billing_Date
		Electric_Bill:   payment_bill.Electric_Bill,   // ตั้งค่าฟิลด์ Electric_Bill
		Water_Bill:      payment_bill.Water_Bill,      // ตั้งค่าฟิลด์ Water_Bill
		Payment_Balance: payment_bill.Payment_Balance, // ตั้งค่าฟิลด์ Payment_Balance
		Employee:        employee,                     // โยงความสัมพันธ์กับ Entity Employee
		Booking:         booking,                      // โยงความสัมพันธ์กับ Entity Booking
		Semester:        semester,                     // โยงความสัมพันธ์กับ Entity Semester
	}

	// 14: บันทึก
	if err := entity.DB().Create(&pb).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": pb})
}

// GET /payment_bill/:id
func GetPayment_Bill(c *gin.Context) {
	var payment_bill entity.Payment_Bill
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&payment_bill); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment_bill not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payment_bill})
}

// GET /payment_bills
func ListPayment_Bills(c *gin.Context) {
	var payment_bills []entity.Payment_Bill
	if err := entity.DB().Raw("SELECT * FROM payment_bills").Find(&payment_bills).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payment_bills})
}

// DELETE /payment_bills/:id
func DeletePayment_Bill(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM payment_bills WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment_bill not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /payment_bills
func UpdatePayment_Bill(c *gin.Context) {
	var payment_bill entity.Payment_Bill
	if err := c.ShouldBindJSON(&payment_bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", payment_bill.ID).First(&payment_bill); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment_bill not found"})
		return
	}

	if err := entity.DB().Save(&payment_bill).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payment_bill})
}
