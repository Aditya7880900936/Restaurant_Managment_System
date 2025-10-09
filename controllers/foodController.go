package controllers

import "github.com/gin-gonic/gin"

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func round(num float64) int {
	return 2
}

func toFixed(num float64, precision int) float64 {
	output := float64(round(num*float64(precision))) / float64(precision)
	return output
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
