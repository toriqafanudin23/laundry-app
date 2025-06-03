package handler

import (
	"database/sql"
	"fmt"
	"laundry-app/entity"
	"log"
	"math/rand"
	"os"
	"time"

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

func GetAllCustomer() []entity.Customer {
	db := connectDb()
	defer db.Close()

	sqlStatement := "SELECT * FROM customer;"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	customers := scanCustomer(rows)

	return customers
}

func scanCustomer(rows *sql.Rows) []entity.Customer {
	customers := []entity.Customer{}
	var err error

	for rows.Next() {
		c := entity.Customer{}
		err := rows.Scan(&c.Customer_id, &c.Name, &c.Phone, &c.Address, &c.Created_at, &c.Updated_at)
		if err != nil {
			panic(err)
		}

		customers = append(customers, c)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return customers
}

func AddCustomer(customer entity.Customer) {
	db := connectDb()
	defer db.Close()
	var err error

	sqlStatement := "INSERT INTO customer (customer_id, name, phone, address, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6);"

	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(900000) + 100000
	customer.Customer_id = id

	_, err = db.Exec(sqlStatement, id, customer.Name, customer.Phone, customer.Address, time.Now(), time.Now())

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Insert Data!")
	}
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
