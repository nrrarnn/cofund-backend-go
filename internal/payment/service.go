package payment

import (
	"time"
)

type PaymentService interface {
	CreateComboPayment(input CreateComboPaymentRequest) error
}

type paymentService struct {
	repo PaymentRepository
}

func NewPaymentService(repo PaymentRepository) PaymentService {
	return &paymentService{repo}
}

func (s *paymentService) CreateComboPayment(input CreateComboPaymentRequest) error {
	payDate, err := time.Parse("2006-01-02", input.PayDate)
	if err != nil {
		return err
	}

	installment := Payment{
		CustomerID: input.CustomerID,
		LoanID:     &input.LoanID,
		Amount:     input.InstallmentAmount,
		Type:       "installment",
		PayDate:    payDate,
		CreatedAt:  time.Now(),
	}

	mandatory := Payment{
		CustomerID: input.CustomerID,
		Amount:     input.MandatoryAmount,
		Type:       "mandatory",
		PayDate:    payDate,
		CreatedAt:  time.Now(),
	}

	if err := s.repo.Create(&installment); err != nil {
		return err
	}

	if err := s.repo.Create(&mandatory); err != nil {
		return err
	}

	return nil
}
