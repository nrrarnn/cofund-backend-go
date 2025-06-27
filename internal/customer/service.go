package customer

import (
	"github.com/nrrarnn/cofund-backend/internal/customer/model"
)

type CustomerService interface {
	CreateCustomer(customer *model.Customer) error
}

type customerService struct {
	repo CustomerRepository
}

func NewCustomerService(repo CustomerRepository) CustomerService {
	return &customerService{repo}
}

func (s *customerService) CreateCustomer(customer *model.Customer) error {
	return s.repo.Create(customer)
}
