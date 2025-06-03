package main

import (
	"laundry-app/handler"
	"laundry-app/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", handler.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware)
	{
		auth.GET("/customers/:sorted", handler.GetAllCustomer)
		auth.POST("/customer", handler.AddCustomer)
		auth.PUT("/customer/:id", handler.UpdateCustomer)
		auth.DELETE("/customer/:id", handler.DeleteCustomer)
	}

	port := ":" + os.Getenv("GIN_PORT")
	r.Run(port)
}
