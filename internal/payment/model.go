package payment

import "time"

type Payment struct {
	ID         uint      `gorm:"primaryKey"`
	CustomerID uint      `gorm:"not null"`
	LoanID     *uint     
	Amount     int       `gorm:"not null"`
	Type       string    `gorm:"not null"` 
	PayDate    time.Time `gorm:"not null"`
	CreatedAt  time.Time
}