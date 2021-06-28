package order

import "time"

const (
	statusAccepted = "Accepted"
)

type Order struct{
	ID           string `json:"orderId"`
	MarketID     string `json:"marketId"`
	Side         string `json:"side"`
	Type         string `json:"type"`
	CreationTime string `json:"creationTime"`
	Price        string `json:"price"`
	Amount       string `json:"amount"`
	OpenAmount   string `json:"openAmount"`
	Status       string `json:"status"`
}

func NewOrder(
	marketID string,
	price string,
	amount string,
	orderType string,
	side string,
) (Order, error) {
	now := time.Now()

	order := Order{
		ID:           "1234",
		MarketID:     marketID,
		Side:         side,
		Type:         orderType,
		CreationTime: now.Format(time.RFC3339Nano),
		Price:        price,
		Amount:       amount,
		OpenAmount:   amount,
		Status:       statusAccepted,
	}

	return order, nil
}
