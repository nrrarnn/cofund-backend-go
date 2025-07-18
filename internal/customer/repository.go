package customer

import (
	"github.com/nrrarnn/cofund-backend/config"
	"github.com/nrrarnn/cofund-backend/internal/customer/model"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
	GetAll() ([]model.Customer, error)
	Update(customer *model.Customer) error
	Delete(customerID uint) error
	FindByID(id uint) (*model.Customer, error)   
}

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepository() CustomerRepository {
	return &customerRepo{db: config.DB}
}

func (r *customerRepo) Update(customer *model.Customer) error {
	return r.db.Save(customer).Error
}

func (r *customerRepo) Delete(customerID uint) error {
	return r.db.Delete(&model.Customer{}, customerID).Error
}
func (r *customerRepo) Create(customer *model.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepo) GetAll() ([]model.Customer, error) {
	var customers []model.Customer
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r *customerRepo) FindByID(id uint) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.First(&customer, id).Error
	return &customer, err
}