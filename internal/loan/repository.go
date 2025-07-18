package loan

import (
	"gorm.io/gorm"
)

type LoanRepository interface {
	Create(loan *Loan) error
	FindByCustomerID(customerID uint) ([]Loan, error)
	Update(loan Loan) error
	FindByID(id uint) (Loan, error)
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

func (r *loanRepository) FindByCustomerID(customerID uint) ([]Loan, error) {
	var loans []Loan
	err := r.db.Where("customer_id = ?", customerID).Find(&loans).Error
	return loans, err
}

func (r *loanRepository) Update(loan Loan) error {
	return r.db.Save(loan).Error
}

func (r *loanRepository) FindByID(id uint) (Loan, error) {
	var loan Loan
	err := r.db.First(&loan, id).Error
	return loan, err
}
