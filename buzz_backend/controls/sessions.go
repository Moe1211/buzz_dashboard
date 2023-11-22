package controls

import (
	"net/http"
	"time"

	"github.com/athunlal/config"
	"github.com/athunlal/models"
	"github.com/gin-gonic/gin"
)

// StartSession handles starting a new session or creating a new one if the room_id is already in use
func StartSession(c *gin.Context) {
	// Extract the room_id from the request or any other method
	var request struct {
		RoomID int `json:"room_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if there is an ongoing session for the specified room_id
	db := config.DB
	var existingSession models.Session
	result := db.Where("room_id = ? AND end_time IS NULL", request.RoomID).First(&existingSession).Error
	if result == nil {
		// A session is already ongoing for this room_id
		c.JSON(http.StatusBadRequest, gin.H{"error": "Session already ongoing for this room_id"})
		return
	}

	// Create a new session instance
	newSession := models.Session{
		UserID:     1, // Replace with the actual user_id // 1 == Guest
		RoomID:     request.RoomID,
		StartTime:  time.Now(),
		EndTime:    time.Time{}, // Initialize to zero value, indicating the session is ongoing
		TotalPrice: 0,           // Initialize to zero value, update as needed
	}

	// Create the new session
	db.Create(&newSession)

	c.JSON(http.StatusOK, gin.H{"message": "Session started successfully"})
}

// CloseSession handles closing a session and moving it to sessions_history
func CloseSession(c *gin.Context) {
	// Extract the room_id from the request or any other method
	var request struct {
		RoomID int `json:"room_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the existing session with end_time
	db := config.DB
	var session models.Session
	result := db.Where("room_id = ? AND end_time IS NULL", request.RoomID).First(&session).Error

	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found or already closed"})
		return
	}

	// Calculate the duration of stay in hours
	startTime := session.StartTime
	endTime := time.Now()
	duration := endTime.Sub(startTime).Hours()

	// Format the end time to your desired format
	// endTimeFormatted := endTime.Format("2006-01-02 15:04:05.000000-07:00")

	// Calculate total price of all items in buffet_order
	var totalPrice float64 // TODO what is this ? there is no model/table for totalPrice
	db.Model(&totalPrice).
		Select("COALESCE(SUM(food_and_drinks.price * buffet_orders.quantity), 0) + ? as total_price", duration*100).
		Joins("JOIN buffet_orders ON buffet_orders.room_id = ?", request.RoomID).
		Joins("JOIN food_and_drinks ON food_and_drinks.item_id = buffet_orders.item_id").
		Scan(&totalPrice)

	// Update the total price and end time in the session
	db.Model(&session).Updates(models.Session{
		TotalPrice: totalPrice,
		EndTime:    endTime,
	})

	// Move the session to sessions_history
	sessionHistory := models.SessionHistory{
		UserID:    session.UserID,
		RoomID:    session.RoomID,
		StartTime: startTime,
		// TODO Fix endTime/Duration
		EndTime:    endTime,
		TotalPrice: totalPrice,
	}
	db.Create(&sessionHistory)

	c.JSON(http.StatusOK, gin.H{"message": "Session closed successfully", "total_price": totalPrice})
}

// ViewSessionsHistory handles viewing sessions history either for all rooms or for a specified room_id
func ViewSessionsHistory(c *gin.Context) {
	// Extract the room_id from the JSON body of the request
	var request struct {
		RoomID int `json:"room_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add your logic to retrieve sessions_history based on the provided room_id
	// If room_id is zero, retrieve sessions_history for all rooms

	// Example logic to retrieve sessions_history
	var sessionsHistory []models.SessionHistory
	db := config.DB

	if request.RoomID == 0 {
		// Retrieve sessions_history for all rooms
		db.Find(&sessionsHistory)
	} else {
		// Retrieve sessions_history for the specified room_id
		db.Where("room_id = ?", request.RoomID).Find(&sessionsHistory)
	}

	// Send the response
	c.JSON(http.StatusOK, gin.H{"sessions_history": sessionsHistory})
}
