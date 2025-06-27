package routes

import (
	"github.com/nrrarnn/cofund-backend/internal/admin"
	"github.com/nrrarnn/cofund-backend/internal/customer"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	repo := admin.NewAdminRepository()
	service := admin.NewAdminService(repo)
	
	service.SeedDefaultAdmin()
	
	handler := admin.NewAdminHandler(service)

	customerRepo := customer.NewCustomerRepository()
	customerService := customer.NewCustomerService(customerRepo)
	customerHandler := customer.NewCustomerHandler(customerService)

	
	api := app.Group("/api")
	api.Post("/login", handler.Login)
	api.Post("/customer", customerHandler.Create)
}