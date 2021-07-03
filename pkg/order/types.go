package order

const (
	statusAccepted = "Accepted"
)

var (
	validOrderTypes = map[string]struct{}{
		"Limit": {},
		"Market": {},
		"Stop Limit": {},
		"Stop": {},
		"Take Profit": {},
	}

	validOrderSides = map[string]struct{}{
		"Bid": {},
		"Ask": {},
	}
)

func isValidOrderType(orderType string) bool {
	_, ok := validOrderTypes[orderType]
	return ok
}

func isValidOrderSide(orderSide string) bool {
	_, ok := validOrderSides[orderSide]
	return ok
}
