package loan

type CreateLoanRequest struct {
	CustomerID uint   `json:"customer_id"` 
	Amount     int    `json:"amount"`      
	ServiceFee int    `json:"service_fee"` 
	Date       string `json:"date"`        
}

type UpdateLoanRequest struct {
	Amount     int     `json:"amount"`
	ServiceFee int     `json:"service_fee"`
	LoanDate   string  `json:"loan_date"`
}