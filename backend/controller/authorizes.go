package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/flugika/project-sa65/entity"
	"github.com/flugika/project-sa65/service"
	"golang.org/x/crypto/bcrypt"
)

// LoginPayload login body
type LoginPayload struct {
	Employee_Email    		string `json:"employee_email"`
	Employee_Personal_ID 	string `json:"employee_personal_id"`
}

// SignUpPayload signup body
type SignUpPayload struct {
	Name     				string `json:"name"`
	Employee_Email    		string `json:"employee_email"`
	Employee_Personal_ID 	string `json:"employee_personal_id"`
}

// LoginResponse token response
type LoginResponse struct {
	Token string `json:"token"`
	ID    uint   `json:"id"`
}

func Login(c *gin.Context) {
	var payload LoginPayload
	var employee entity.Employee

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา user ด้วย email ที่ผู้ใช้กรอกเข้ามา
	if err := entity.DB().Raw("SELECT * FROM employees WHERE employee_email = ?", payload.Employee_Email).Scan(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(employee.Employee_Personal_ID), []byte(payload.Employee_Personal_ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is incerrect"})
		return
	}

	// กำหนดค่า SecretKey, Issuer และระยะเวลาหมดอายุของ Token สามารถกำหนดเองได้
	// SecretKey ใช้สำหรับการ sign ข้อความเพื่อบอกว่าข้อความมาจากตัวเราแน่นอน
	// Issuer เป็น unique id ที่เอาไว้ระบุตัว client
	// ExpirationHours เป็นเวลาหมดอายุของ token

	jwtWrapper := service.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(employee.Employee_Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginResponse{
		Token: signedToken,
		ID:    employee.ID,
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}

// POST /create
func CreateEmployee(c *gin.Context) {
	var payload SignUpPayload
	var employee entity.Employee

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Employee_Personal_ID), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}

	employee.Employee_Name = payload.Name
	employee.Employee_Email = payload.Employee_Email
	employee.Employee_Personal_ID = string(hashPassword)

	if err := entity.DB().Create(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": employee})
}