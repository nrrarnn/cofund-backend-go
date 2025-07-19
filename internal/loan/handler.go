package loan

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
)

type LoanHandler struct {
	service LoanService
}

func NewLoanHandler(service LoanService) *LoanHandler {
	return &LoanHandler{service}
}

func (h *LoanHandler) CreateLoan(c *fiber.Ctx) error {
	var input CreateLoanRequest

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	if err := h.service.CreateLoan(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create loan",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Loan created successfully",
	})
}

func (h *LoanHandler) GetLoansByCustomerID(c *fiber.Ctx) error {
	idParam := c.Params("customer_id")

	customerID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid customer ID",
		})
	}

	loans, err := h.service.GetLoansByCustomerID(uint(customerID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get loans",
		})
	}

	return c.JSON(loans)
}

func (h *LoanHandler) UpdateLoan(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	var req UpdateLoanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := h.service.UpdateLoan(uint(id), req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update loan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Loan updated successfully",
	})
}

func (h *LoanHandler) DeleteLoan(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	err = h.service.DeleteLoan(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete loan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Loan deleted successfully",
	})
}

