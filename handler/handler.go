package handler

import (
	"database/sql"
	"fmt"
	"laundry-app/entity"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var psqlInfo string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func GetAllCustomer(c *gin.Context) {
	db := connectDb()
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

	c.JSON(http.StatusOK, customers)
}

func AddCustomer(c *gin.Context) {
	db := connectDb()
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

	c.JSON(http.StatusCreated, s)
}

func UpdateCustomer(c *gin.Context) {
	db := connectDb()
	defer db.Close()

	id := c.Param("id")
	var s entity.Customer
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "UPDATE customer SET name = $2, phone = $3, address = $4, updated_at = $5 WHERE customer_id = $1"
	_, err := db.Exec(query, id, s.Name, s.Phone, s.Address, time.Now())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, s)
}

func DeleteCustomer(c *gin.Context) {
	db := connectDb()
	defer db.Close()
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM customer WHERE customer_id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succesfully Deleted Customer!"})
}

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully connected!")
	}
	return db
}
