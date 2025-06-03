package handler

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtKey []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not loaded")
	}

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Fatal("SECRET_KEY is not set in environment")
	}
	jwtKey = []byte(secret)
}

func Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Dummy auth, ganti dengan query ke database
	if credentials.Username != "admin" || credentials.Password != "1234" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": credentials.Username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
