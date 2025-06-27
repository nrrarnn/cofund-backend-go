package main

import (
	"github.com/nrrarnn/cofund-backend/config"
	"github.com/nrrarnn/cofund-backend/routes"
	"github.com/nrrarnn/cofund-backend/internal/admin/model"
	"github.com/gofiber/fiber/v2"
	"os"
	customerModel "github.com/nrrarnn/cofund-backend/internal/customer/model"
)

func main() {
	app := fiber.New()

	config.InitDB()

	config.DB.AutoMigrate(&model.Admin{})
	config.DB.AutoMigrate(&customerModel.Customer{})

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
