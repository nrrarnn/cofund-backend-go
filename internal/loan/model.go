package loan

import "time"

type Loan struct {
	ID         uint      `gorm:"primaryKey"`
	CustomerID uint      `gorm:"not null"`
	Amount     int       `gorm:"not null"`      
	ServiceFee int       `gorm:"not null"`      
	Total      int       `gorm:"not null"`      
	Status     string    `gorm:"default:'active'"`
	LoanDate   time.Time `gorm:"not null"`
	CreatedAt  time.Time
}
