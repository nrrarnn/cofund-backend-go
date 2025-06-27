package admin

import (
	"github.com/nrrarnn/cofund-backend/internal/admin/model"
	"github.com/nrrarnn/cofund-backend/pkg/utils"
	"errors"
	"fmt"
)

type AdminService interface {
	Login(username, password string) (string, error)
	SeedDefaultAdmin()
}

type adminService struct {
	repo AdminRepository
}

func NewAdminService(repo AdminRepository) AdminService {
	return &adminService{repo: repo}
}

func (s *adminService) Login(username, password string) (string, error) {
	admin, err := s.repo.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, admin.Password) {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(admin.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *adminService) SeedDefaultAdmin() {
	admin, err := s.repo.FindByUsername("admin")
	if err == nil && admin != nil {
		fmt.Println("⚠️ Admin default sudah ada")
		return
	}

	fmt.Println("🌱 Membuat admin default...")

	hashed := utils.HashPassword("admin123")

	s.repo.Create(&model.Admin{
		Username: "admin",
		Password: hashed,
	})
}
