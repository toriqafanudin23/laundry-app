package main

import (
	"laundry-app/handler"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/customers/:sorted", handler.GetAllCustomer)
	r.POST("/customer", handler.AddCustomer)
	r.PUT("/customer/:id", handler.UpdateCustomer)
	r.DELETE("/customer/:id", handler.DeleteCustomer)

	port := ":" + os.Getenv("GIN_PORT")
	r.Run(port)
}
