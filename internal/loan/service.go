package loan

import (
	"time"

)

type LoanService interface {
	CreateLoan(input CreateLoanRequest) error
	GetLoansByCustomerID(customerID uint) ([]Loan, error)
}

type loanService struct {
	repo LoanRepository
}

func NewLoanService(repo LoanRepository) LoanService {
	return &loanService{repo}
}

func (s *loanService) CreateLoan(input CreateLoanRequest) error {
	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		return err
	}

	loan := Loan{
		CustomerID: input.CustomerID,
		Amount:     input.Amount,
		ServiceFee: input.ServiceFee,
		Total:      input.Amount + input.ServiceFee,
		Status:     "active",
		LoanDate:   date,
		CreatedAt:  time.Now(),
	}

	return s.repo.Create(&loan)
}

func (s *loanService) GetLoansByCustomerID(customerID uint) ([]Loan, error) {
	return s.repo.FindByCustomerID(customerID)
}
