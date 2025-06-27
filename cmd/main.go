package main

import (
	"github.com/nrrarnn/cofund-backend/config"
	"github.com/nrrarnn/cofund-backend/routes"
	"github.com/nrrarnn/cofund-backend/internal/admin/model"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	app := fiber.New()

	config.InitDB()

	config.DB.AutoMigrate(&model.Admin{})

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
