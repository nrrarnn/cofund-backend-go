package routes

import (
	"github.com/nrrarnn/cofund-backend/internal/admin"
	"github.com/nrrarnn/cofund-backend/internal/customer"
	"github.com/nrrarnn/cofund-backend/internal/loan"
	"github.com/nrrarnn/cofund-backend/internal/payment"
	"github.com/nrrarnn/cofund-backend/config"
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

	loanRepo := loan.NewLoanRepository(config.DB)
	loanService := loan.NewLoanService(loanRepo)
	loanHandler := loan.NewLoanHandler(loanService)

	paymentRepo := payment.NewPaymentRepository(config.DB)
	paymentService := payment.NewPaymentService(paymentRepo)	
	paymentHandler := payment.NewPaymentHandler(paymentService)	
	
	api := app.Group("/api")
	api.Post("/login", handler.Login)
	api.Post("/customer", customerHandler.Create)
	api.Post("/loan", loanHandler.CreateLoan)
	api.Post("/payment", paymentHandler.CreateComboPayment)
	api.Get("/customers", customerHandler.GetAllCustomers)
	api.Put("/customers/:id", customerHandler.UpdateCustomer)
	api.Delete("/customers/:id", customerHandler.DeleteCustomer)
	api.Get("/loans/:customer_id", loanHandler.GetLoansByCustomerID)
	api.Put("/loans/:id", loanHandler.UpdateLoan)
	api.Delete("/loans/:id", loanHandler.DeleteLoan)
	api.Put("/payments/:id", paymentHandler.UpdatePayment)
	api.Delete("/payments/:id", paymentHandler.DeletePayment)
}