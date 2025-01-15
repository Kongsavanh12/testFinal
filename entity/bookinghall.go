package entity

import (
	"gorm.io/gorm"
)
type BookingHall struct {
	gorm.Model
	CustomerName	string	`valid: "required~CustomerName is required"`
	CustomerPhone	string	`valid: "required~CustomerPhone is required"`
	CustomerEmail	string	`valid:"required~CustomerEmail is required"`
}