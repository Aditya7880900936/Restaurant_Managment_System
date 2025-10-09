package controllers

import "github.com/gin-gonic/gin"

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Get all users route"})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Get user route"})
	}
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Signup route"})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Login route"})
	}
}

func HashPassword(password string) string {
	return password
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	return userPassword == providedPassword, "success"
}
