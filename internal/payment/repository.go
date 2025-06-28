package payment

import (
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *Payment) error
	FindAll() ([]Payment, error)
}

type paymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepo{db: db}
}

func (r *paymentRepo) Create(payment *Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepo) FindAll() ([]Payment, error) {
	var payments []Payment
	err := r.db.Find(&payments).Error
	return payments, err
}
