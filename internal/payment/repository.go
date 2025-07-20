package payment

import (
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *Payment) error
	FindAll() ([]Payment, error)
	Update(payment *Payment) error
	Delete(id uint) error
	FindByID(id uint) (*Payment, error)
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

func (r *paymentRepo) Update(payment *Payment) error {
	return r.db.Save(payment).Error
}

func (r *paymentRepo) Delete(id uint) error {
	return r.db.Delete(&Payment{}, id).Error
}

func (r *paymentRepo) FindByID(id uint) (*Payment, error) {
	var payment Payment
	err := r.db.First(&payment, id).Error
	return &payment, err
}
