package customer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nrrarnn/cofund-backend/internal/customer/model"
)

type CustomerHandler struct {
	service CustomerService
}

func NewCustomerHandler(service CustomerService) *CustomerHandler {
	return &CustomerHandler{service}
}

func (h *CustomerHandler) Create(c *fiber.Ctx) error {
	var customer model.Customer
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.service.CreateCustomer(&customer); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create customer"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Customer created successfully",
		"data":    customer,
	})
}
