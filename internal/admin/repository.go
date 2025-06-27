package admin

import (
	"gorm.io/gorm"
	"github.com/nrrarnn/cofund-backend/config"
	"github.com/nrrarnn/cofund-backend/internal/admin/model"
)

type AdminRepository interface {
	FindByUsername(username string) (*model.Admin, error)
	Create(admin *model.Admin) error
}

type adminRepo struct {
	db *gorm.DB
}

func NewAdminRepository() AdminRepository {
	return &adminRepo{db: config.DB}
}

func (r *adminRepo) FindByUsername(username string) (*model.Admin, error) {
	var admin model.Admin
	err := r.db.Where("username = ?", username).First(&admin).Error
	return &admin, err
}

func (r *adminRepo) Create(admin *model.Admin) error {
	return r.db.Create(admin).Error
}
