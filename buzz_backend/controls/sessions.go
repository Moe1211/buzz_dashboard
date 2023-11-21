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

	// Create a new session instance
	newSession := models.Session{
		UserID:     1, // Replace with the actual user_id
		RoomID:     request.RoomID,
		StartTime:  time.Now(),
		EndTime:    time.Time{}, // Initialize to zero value, indicating the session is ongoing
		TotalPrice: 0,           // Initialize to zero value, update as needed
	}

	db := config.DB
	result := db.First(&newSession, "room_id = ?", newSession.RoomID).Error
	if result != nil {
		db.Create(&newSession)
	} else {
		db.Model(&newSession).Where("room_id = ?", newSession.RoomID).Update("start_time", time.Now())
	}

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

	if result != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found or already closed"})
		return
	}

	db.Model(&session).Update("end_time", time.Now())

	// Move the session to sessions_history
	sessionHistory := models.SessionHistory{
		UserID:     session.UserID,
		RoomID:     session.RoomID,
		StartTime:  session.StartTime,
		EndTime:    session.EndTime,
		TotalPrice: session.TotalPrice,
	}
	db.Create(&sessionHistory)

	// Calculate total price of all items in buffet_order
	var totalPrice float64
	db.Model(&sessionHistory).Select("COALESCE(SUM(price * quantity), 0) as total_price").Joins("JOIN buffet_orders ON buffet_orders.room_id = session_history.room_id").Joins("JOIN food_and_drinks ON food_and_drinks.item_id = buffet_orders.item_id").Scan(&totalPrice)

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
