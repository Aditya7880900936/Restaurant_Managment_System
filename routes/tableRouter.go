package routes

import (
	"github.com/Aditya7880900936/Restaurant_Managment_System/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controllers.GetTables())
	incomingRoutes.GET("/tables/:table_id", controllers.GetTable())
	incomingRoutes.POST("/tables", controllers.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id", controllers.UpdateTable())
}