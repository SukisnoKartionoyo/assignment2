package main

import (
	"sesi7/config"
	"sesi7/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/orders", controllers.CreateOrder)
	r.GET("/order/:id", controllers.GetOrderByID)
	r.PUT("/order/:id", controllers.UpdateOrderByID)
	r.DELETE("/order/:id", controllers.DeleteOrderById)
	config.ConnectDatabase()

	r.Run()
}
