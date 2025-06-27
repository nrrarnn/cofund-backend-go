package main

import (
	"github.com/nrrarnn/cofund-backend/config"
	"github.com/nrrarnn/cofund-backend/routes"
	"github.com/nrrarnn/cofund-backend/internal/admin/model"
	"github.com/gofiber/fiber/v2"
	"os"
	loan "github.com/nrrarnn/cofund-backend/internal/loan"
	customerModel "github.com/nrrarnn/cofund-backend/internal/customer/model"
)

func main() {
	app := fiber.New()

	config.InitDB()

	config.DB.AutoMigrate(&model.Admin{})
	config.DB.AutoMigrate(&customerModel.Customer{})
	config.DB.AutoMigrate(&loan.Loan{})

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
