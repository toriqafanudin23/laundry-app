package handler

import (
	"laundry-app/connectdb"
	"laundry-app/entity"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func GetAllCustomer(c *gin.Context) {
	db := connectdb.ConnectDb()
	defer db.Close()
	sorted := c.Param("sorted")

	query := "SELECT * FROM customer ORDER BY " + sorted + " ASC;"
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var customers []entity.Customer

	for rows.Next() {
		var s entity.Customer
		if err := rows.Scan(&s.Customer_id, &s.Name, &s.Phone, &s.Address, &s.Created_at, &s.Updated_at); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		customers = append(customers, s)
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    customers,
		"message": "Succesfully Fetch Customers",
	})
}

func AddCustomer(c *gin.Context) {
	db := connectdb.ConnectDb()
	defer db.Close()

	var s entity.Customer
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(900000) + 100000

	query := "INSERT INTO customer (customer_id, name, phone, address, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6);"
	_, err := db.Exec(query, id, s.Name, s.Phone, s.Address, time.Now(), time.Now())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    s,
		"message": "Succesfully Insert Data",
	})
}

func UpdateCustomer(c *gin.Context) {
	db := connectdb.ConnectDb()
	defer db.Close()

	id := c.Param("id")
	var s entity.Customer
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.Customer_id, _ = strconv.Atoi(id)
	s.Updated_at = time.Now()
	query := "UPDATE customer SET name = $2, phone = $3, address = $4, updated_at = $5 WHERE customer_id = $1"
	_, err := db.Exec(query, s.Customer_id, s.Name, s.Phone, s.Address, s.Updated_at)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    s,
		"message": "Succesfully Update Data!",
	})
}

func DeleteCustomer(c *gin.Context) {
	db := connectdb.ConnectDb()
	defer db.Close()
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM customer WHERE customer_id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succesfully Deleted Customer!"})
}
