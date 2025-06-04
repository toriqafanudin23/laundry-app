package main

import (
	"laundry-app/handler"
	// "laundry-app/middleware"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Tambahkan ini
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // atau spesifik: "http://localhost:5500"
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/customer", handler.AddCustomer)
	r.PUT("/customer/:id", handler.UpdateCustomer)
	r.DELETE("/customer/:id", handler.DeleteCustomer)
	r.GET("/customers/:sorted", handler.GetAllCustomer)

	port := ":" + os.Getenv("GIN_PORT")
	r.Run(port)
}

// r.POST("/login", handler.Login)

// auth := r.Group("/")
// auth.Use(middleware.AuthMiddleware)
// {

// }
