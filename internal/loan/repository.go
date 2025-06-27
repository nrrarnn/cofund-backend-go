package loan

import (
	"gorm.io/gorm"
)

type LoanRepository interface {
	Create(loan *Loan) error
}

type loanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepository{db: db}
}

func (r *loanRepository) Create(loan *Loan) error {
	return r.db.Create(loan).Error
}