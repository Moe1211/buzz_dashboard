package routes

import (
	"net/http"

	"github.com/athunlal/controls"
	"github.com/gin-gonic/gin"

	"time"
)

func UserRoutes(c *gin.Engine) {
	user := c.Group("/user")
	{
		// Existing routes
		user.POST("/startsession", controls.StartSession)
		user.POST("/closesession", controls.CloseSession)

		user.GET("/returnbuffetorder", controls.ReturnBuffetOrder)
		user.POST("/addfoodanddrink", controls.AddBeverageToBuffetOrder)
		user.POST("/addbeverage", controls.AddBeverageToBuffetOrder)

		// Add a new route for viewing sessions_history either for all rooms or a specified room_id
		user.POST("/viewsessionshistory", controls.ViewSessionsHistory)

		// Invoice routes
		// User.GET("/invoice", middlereware.UserAuth, controls.InvoiceF)
		// User.GET("/invoice/download", middlereware.UserAuth, controls.Download)
	}
}

// User represents the user model
type User struct {
	UserID          int       `json:"user_id"`
	Name            string    `json:"name"`
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	StartedAt       time.Time `json:"started_at"`
	Duration        int       `json:"duration"`
	CurrentDuration int       `json:"current_duration"`
	PricePerHour    float64   `json:"price_per_hour"`
	TotalTime       int       `json:"total_time"`
	TotalBuffet     float64   `json:"total_buffet"`
	FinalCost       float64   `json:"final_cost"`
}

// Room represents the room model
type Room struct {
	RoomID          int       `json:"room_id"`
	RoomNumber      int       `json:"room_number"`
	CustomerID      int       `json:"customer_id"`
	Game            string    `json:"game"`
	StartedAt       time.Time `json:"started_at"`
	Duration        int       `json:"duration"`
	CurrentDuration int       `json:"current_duration"`
	PricePerHour    float64   `json:"price_per_hour"`
	TotalTime       int       `json:"total_time"`
	TotalBuffet     float64   `json:"total_buffet"`
	FinalCost       float64   `json:"final_cost"`
}

// FoodAndDrink represents the food and drink model
type FoodAndDrink struct {
	ItemID   int     `json:"item_id"`
	ItemName string  `json:"item_name"`
	Price    float64 `json:"price"`
}

// Session represents the session model
type Session struct {
	SessionID  int       `json:"session_id"`
	UserID     int       `json:"user_id"`
	RoomID     int       `json:"room_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	TotalPrice float64   `json:"total_price"`
}

// SessionHistory represents the session history model
type SessionHistory struct {
	HistoryID  int       `json:"history_id"`
	UserID     int       `json:"user_id"`
	RoomID     int       `json:"room_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	TotalPrice float64   `json:"total_price"`
}

// BuffetOrder represents the buffet order model
type BuffetOrder struct {
	OrderID  int `json:"order_id"`
	RoomID   int `json:"room_id"`
	ItemID   int `json:"item_id"`
	Quantity int `json:"quantity"`
}

// UserSignUp handles user signup
func UserSignUp(c *gin.Context) {
	// Implement user signup logic
	c.JSON(http.StatusOK, gin.H{"message": "User signed up successfully"})
}

// UserLogin handles user login
func UserLogin(c *gin.Context) {
	// Implement user login logic
	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
}

// AddToCart handles adding items to the cart
func AddToCart(c *gin.Context) {
	// Implement add to cart logic
	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

// ViewCart handles viewing the user's cart
func ViewCart(c *gin.Context) {
	// Implement view cart logic
	c.JSON(http.StatusOK, gin.H{"message": "Viewing user's cart"})
}

// Checkout handles the checkout process
func Checkout(c *gin.Context) {
	// Implement checkout logic
	c.JSON(http.StatusOK, gin.H{"message": "Checkout successful"})
}

// Other functions can be added based on your application's requirements.
