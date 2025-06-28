package payment

import (
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

