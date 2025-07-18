package customer

import (
	"github.com/nrrarnn/cofund-backend/internal/customer/model"
)

type CustomerService interface {
	CreateCustomer(customer *model.Customer) error
	GetAllCustomers() ([]model.Customer, error)
	UpdateCustomer(id uint, input UpdateCustomerRequest) error
	DeleteCustomer(id uint) error

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

func (s *customerService) GetAllCustomers() ([]model.Customer, error) {
	return s.repo.GetAll()
}

func (s *customerService) UpdateCustomer(id uint, input UpdateCustomerRequest) error {
	customer, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	customer.Name = input.Name
	customer.Phone = input.Phone
	customer.Address = input.Address

	return s.repo.Update(customer)
}

func (s *customerService) DeleteCustomer(id uint) error {
	return s.repo.Delete(id)
}
