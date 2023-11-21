package controls

import (
	"net/http"

	"github.com/athunlal/config"
	"github.com/athunlal/models"
	"github.com/gin-gonic/gin"
)

// AddBeverageToBuffetOrder handles adding a beverage to buffet_orders
func AddBeverageToBuffetOrder(c *gin.Context) {
	// Extract the buffet_item to be added from the request
	var request struct {
		BuffetItemID  int `json:"buffet_item_id"`
		BuffetItemQty int `json:"buffet_item_qty"`
		RoomID        int `json:"room_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use room_id to get the corresponding buffet_order
	var buffetOrder models.BuffetOrder
	db := config.DB
	result := db.Where("room_id = ?", request.RoomID).First(&buffetOrder).Error

	defaultQuantity := 1 // Change this to your desired default quantity

	if result != nil {
		// Create a new buffet_order
		buffetOrder = models.BuffetOrder{
			RoomID:   request.RoomID,
			ItemID:   request.BuffetItemID,
			Quantity: defaultQuantity,
		}
		db.Create(&buffetOrder)
	} else {
		// Update the existing buffet_order
		quantityToAdd := defaultQuantity
		if request.BuffetItemQty != 0 {
			quantityToAdd = request.BuffetItemQty
		}
		db.Model(&buffetOrder).Update("quantity", buffetOrder.Quantity+quantityToAdd)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Beverage added to Buffet Orders"})
}

// ReturnBuffetOrder handles returning the buffet_order corresponding to the correct session
func ReturnBuffetOrder(c *gin.Context) {
	// Extract the room_id from the JSON body of the request
	var request struct {
		RoomID int `json:"room_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use room_id to get the corresponding buffet_order
	var buffetOrder models.BuffetOrder
	db := config.DB
	result := db.Where("room_id = ?", request.RoomID).First(&buffetOrder).Error

	if result != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Buffet Order not found"})
		return
	}

	// Calculate total price of all items in buffet_order
	var totalPrice float64
	db.Model(&buffetOrder).Select("COALESCE(SUM(price * quantity), 0) as total_price").Joins("JOIN food_and_drinks ON food_and_drinks.item_id = buffet_orders.item_id").Scan(&totalPrice)

	c.JSON(http.StatusOK, gin.H{
		"room_id":     buffetOrder.RoomID,
		"item_id":     buffetOrder.ItemID,
		"quantity":    buffetOrder.Quantity,
		"total_price": totalPrice,
	})
}
