package main

import (
	"laundry-app/handler"
	"laundry-app/middleware"
	"os"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/login", handler.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware)
	{
		auth.POST("/customer", handler.AddCustomer)
		auth.PUT("/customer/:id", handler.UpdateCustomer)
		auth.DELETE("/customer/:id", handler.DeleteCustomer)
		auth.GET("/customers/:sorted", handler.GetAllCustomer)
	}

	port := ":" + os.Getenv("GIN_PORT")
	r.Run(port)
}
