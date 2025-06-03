package main

import (
	"laundry-app/entity"
	"laundry-app/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/customers", func(c *gin.Context) {
		customers, err := handler.GetAllCustomer()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    customers,
			"message": "Successfully fetched customers",
		})
	})

	r.POST("/customers", func(c *gin.Context) {
		var newCustomer entity.Customer
		if err := c.BindJSON(&newCustomer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		handler.AddCustomer(newCustomer)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Successfully Insert Data!",
		})

	})

	r.Run(":8080")
}
