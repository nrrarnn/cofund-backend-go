package model

type Customer struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Phone     string `gorm:"unique;not null"`
	Address   string
}
