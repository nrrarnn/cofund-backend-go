package customer

import (
	"strconv"

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

func (h *CustomerHandler) GetAllCustomers(c *fiber.Ctx) error {
	customers, err := h.service.GetAllCustomers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve customers",
		})
	}
	return c.JSON(customers)
}

func (h *CustomerHandler) UpdateCustomer(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID"})
	}

	var req UpdateCustomerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}

	err = h.service.UpdateCustomer(uint(id), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update customer"})
	}

	return c.JSON(fiber.Map{"message": "Customer updated successfully"})
}

func (h *CustomerHandler) DeleteCustomer(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID"})
	}

	err = h.service.DeleteCustomer(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete customer"})
	}

	return c.JSON(fiber.Map{"message": "Customer deleted successfully"})
}


