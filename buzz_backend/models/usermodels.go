package models

import (
	"encoding/json"
	"time"
)

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

// / ToMap converts any object to a map[string]interface{}
func ToMap(obj interface{}) map[string]interface{} {
	var result map[string]interface{}
	JSONData, _ := json.Marshal(obj)
	json.Unmarshal(JSONData, &result)
	return result
}
