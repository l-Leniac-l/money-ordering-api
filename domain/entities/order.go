package entities

type Order struct {
	OrderId        string  `json:"orderId,omitempty"`
	BankId         string  `json:"bankId"`
	Amount         float64 `json:"amount"`
	AdditionalInfo string  `json:"additionalInfo"`
}
