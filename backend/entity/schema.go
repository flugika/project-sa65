package entity

import (
	"time"

	"gorm.io/gorm"
)

// -------------------------------------------------------------------------------------------------
// entity Room_Price สร้างเพื่อทดสอบการคำนวณ
type Room_Price struct {
	gorm.Model
	Room_Price int	`gorm:"references:id"`

	// 1 Room_Price เป็นเจ้าของได้หลาย Room
	Room []Room `gorm:"foreignKey:Room_PriceID"`
}

// entity Room, Student สร้างเพื่อทดสอบ foreign key
type Room struct {
	gorm.Model

	// 1 Room เป็นเจ้าของได้หลาย Booking
	Booking []Booking `gorm:"foreignKey:RoomID"`

	// Room_PriceID ทำหน้าที่เป็น FK
	Room_PriceID int
	Room_Price   Room_Price `gorm:"references:id"`
}

type Student struct {
	gorm.Model
	Student_Number string

	// 1 Student เป็นเจ้าของได้หลาย Booking
	Booking []Booking `gorm:"foreignKey:StudentID"`
}

// -------------------------------------------------------------------------------------------------

type Employee struct {
	gorm.Model

	Employee_Personal_ID string	`json:"-"`
	Employee_Name        string
	Employee_Email       string `gorm:"uniqueIndex"`
	// JobPositionID	int
	// JobPosition   	JobPosition	`gorm:"references:id"`
	// ProvinceID		int
	// Province   		Province	`gorm:"references:id"`
	// GenderID			int
	// Gender			Gender		`gorm:"references:id"`

	// 1 Employee เป็นเจ้าของได้หลาย Payment_Bill
	Payment_Bill []Payment_Bill `gorm:"foreignKey:EmployeeID"`
}

type Semester struct {
	gorm.Model

	Semester string
	Payment_Bill []Payment_Bill `gorm:"foreignKey:SemesterID"`
}

type Booking struct {
	gorm.Model
	Check_In_Date time.Time //datetime

	// RoomID ทำหน้าที่เป็น FK
	RoomID int
	Room   Room	`gorm:"references:id"`

	// StudentID ทำหน้าที่เป็น FK
	StudentID int
	Student   Student `gorm:"references:id"`

	// TimeID int
	// Time   Time	`gorm:"references:id"`

	// 1 Booking เป็นเจ้าของได้หลาย Payment_Bill
	Payment_Bill []Payment_Bill `gorm:"foreignKey:BookingID"`
}

type Payment_Bill struct {
	gorm.Model
	Billing_Date    time.Time
	Electric_Bill float32
	Water_Bill float32
	Payment_Balance float32

	// Employee ทำหน้าที่เป็น FK
	EmployeeID int
	Employee   Employee	`gorm:"references:id"`

	// BookingID ทำหน้าที่เป็น FK
	BookingID int
	Booking   Booking `gorm:"references:id"`
	
	// Semester ทำหน้าที่เป็น FK
	SemesterID int
	Semester   Semester	`gorm:"references:id"`
}
