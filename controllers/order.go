package controllers

import (
	"fmt"
	"net/http"
	"sesi7/config"
	"sesi7/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(c *gin.Context) {
	var order models.Order

	c.ShouldBindJSON(&order)
	fmt.Println(&order)
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't create!"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")

	var order models.Order
	if err := config.DB.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func UpdateOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var order models.Order

	if err := config.DB.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	} else {
		fmt.Println(&order)
		c.ShouldBindJSON(&order)

		if err := config.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&order).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't update!"})
			return
		}

		c.JSON(http.StatusOK, order)
	}
}

func DeleteOrderById(c *gin.Context) {
	id := c.Params.ByName("id")
	var order models.Order
	var item []models.Item

	if err := config.DB.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	fmt.Println(order.Items[0].LineItemID)
	if err := config.DB.Delete(item, "order_id = ?", id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't found!"})
		return
	}

	if err := config.DB.Delete(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't found!"})
		return
	}

	c.JSON(http.StatusOK, order)
}
