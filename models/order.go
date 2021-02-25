package models

import (
	"time"
)

//Item itu buat data item
type Item struct {
	// gorm.Model
	LineItemID  uint   `json:"lineItemId" gorm:"primaryKey"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"-"`
}

//Order itu sesuatu banget
type Order struct {
	// gorm.Model
	OrderID      uint      `json:"orderId" gorm:"primaryKey"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID"`
}
