package main

import (
	"os"

	"github.com/Aditya7880900936/Restaurant_Managment_System/database"
	"github.com/Aditya7880900936/Restaurant_Managment_System/middlewares"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
    router.UserRoutes(router)
	router.Use(middlewares.Authentication())
	router.FoodRoutes(router)
	router.InvoiceRoutes(router)
	router.MenuRoutes(router)
	router.OrderItemRoutes(router)
	router.OrderRoutes(router)
	router.TableRoutes(router)

	router.Run(":" + port)
}
