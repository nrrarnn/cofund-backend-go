package admin

import (
	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct {
	service AdminService
}

func NewAdminHandler(service AdminService) *AdminHandler {
	return &AdminHandler{service}
}

func (h *AdminHandler) Login(c *fiber.Ctx) error {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	token, err := h.service.Login(input.Username, input.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": token})
}
