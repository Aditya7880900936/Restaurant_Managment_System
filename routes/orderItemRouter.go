package routes

import (
	"github.com/Aditya7880900936/Restaurant_Managment_System/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())
}