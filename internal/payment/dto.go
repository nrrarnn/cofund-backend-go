package payment

type CreateComboPaymentRequest struct {
	CustomerID        uint   `json:"customer_id"`
	LoanID            uint   `json:"loan_id"`
	InstallmentAmount int    `json:"installment_amount"`
	MandatoryAmount   int    `json:"mandatory_amount"`
	PayDate           string `json:"pay_date"` 
}

type UpdatePaymentRequest struct {
	Amount  int     `json:"amount"`
	Type    string  `json:"type"`     
	PayDate string  `json:"pay_date"` 
}
