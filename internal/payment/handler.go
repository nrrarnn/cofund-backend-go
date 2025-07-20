package payment

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
)

type PaymentHandler struct {
	service PaymentService
}

func NewPaymentHandler(service PaymentService) *PaymentHandler {
	return &PaymentHandler{service}
}

func (h *PaymentHandler) CreateComboPayment(c *fiber.Ctx) error {
	var input CreateComboPaymentRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.service.CreateComboPayment(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to process payment"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Payment successful"})
}

func (h *PaymentHandler) UpdatePayment(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID"})
	}

	var req UpdatePaymentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}

	if err := h.service.UpdatePayment(uint(id), req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update payment"})
	}

	return c.JSON(fiber.Map{"message": "Payment updated successfully"})
}

func (h *PaymentHandler) DeletePayment(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID"})
	}

	if err := h.service.DeletePayment(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete payment"})
	}

	return c.JSON(fiber.Map{"message": "Payment deleted successfully"})
}
