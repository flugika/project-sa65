package main

import (
	"github.com/gin-gonic/gin"
	"github.com/flugika/project-sa65/controller"
	"github.com/flugika/project-sa65/entity"
	"github.com/flugika/project-sa65/middlewares"
)

const PORT = "8080"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())
		{
			// Room_Price Routes
			router.POST("/room_prices", controller.CreateRoom_Price)
			router.GET("/room_price/:id", controller.GetRoom_Price)
			router.GET("/room_prices", controller.ListRoom_Prices)
			router.DELETE("/room_prices/:id", controller.DeleteRoom_Price)
			router.PATCH("/room_prices", controller.UpdateRoom_Price)
			
			// Room Routes
			router.POST("/rooms", controller.CreateRoom)
			router.GET("/room/:id", controller.GetRoom)
			router.GET("/rooms", controller.ListRooms)
			router.DELETE("/rooms/:id", controller.DeleteRoom)
			router.PATCH("/rooms", controller.UpdateRoom)

			// Student Routes
			router.POST("/students", controller.CreateStudent)
			router.GET("/student/:id", controller.GetStudent)
			router.GET("/students", controller.ListStudents)
			router.DELETE("/students/:id", controller.DeleteStudent)
			router.PATCH("/students", controller.UpdateStudent)

			// Employee Routes
			router.GET("/employee/:id", controller.GetEmployee)
			router.GET("/employees", controller.ListEmployees)
			router.DELETE("/employees/:id", controller.DeleteEmployee)
			router.PATCH("/employees", controller.UpdateEmployee)

			// Booking Routes
			router.POST("/bookings", controller.CreateBooking)
			router.GET("/booking/:id", controller.GetBooking)
			router.GET("/bookings", controller.ListBookings)
			router.DELETE("/bookings/:id", controller.DeleteBooking)
			router.PATCH("/bookings", controller.UpdateBooking)

			// Semester Routes
			router.POST("/semesters", controller.CreateSemester)
			router.GET("/semester/:id", controller.GetSemester)
			router.GET("/semesters", controller.ListSemesters)
			router.DELETE("/semesters/:id", controller.DeleteSemester)
			router.PATCH("/semesters", controller.UpdateSemester)

			// Payment_Bill Routes
			router.POST("/payment_bills", controller.CreatePayment_Bill)
			router.GET("/payment_bills", controller.ListPayment_Bills)
			router.GET("/payment_bill/:id", controller.GetPayment_Bill)
			router.PATCH("/payment_bills", controller.UpdatePayment_Bill)
			router.DELETE("/payment_bills/:id", controller.DeletePayment_Bill)
		}
	}

	// Signup User Route
	r.POST("/signup", controller.CreateEmployee)
	// login User Route
	r.POST("/login", controller.Login)

	// Run the server go run main.go
	r.Run("localhost: " + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}