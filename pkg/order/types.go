package order

const (
	statusAccepted           = "Accepted"
	statusPlaced             = "Placed"
	statusPartiallyMatched   = "Partially Matched"
	statusFullyMatched       = "Fully Matched"
	statusCancelled          = "Cancelled"
	statusPartiallyCancelled = "Partially Cancelled"
	statusFailed             = "Failed"

	typeLimit      = "Limit"
	typeMarket     = "Market"
	typeStopLimit  = "Stop Limit"
	typeStop       = "Stop"
	typeTakeProfit = "Take Profit"

	sideBid = "Bid"
	sideAsk = "Ask"

	tradeLiquidityTypeMaker = "Maker"
	tradeLiquidityTypeTaker = "Taker"
)

var (
	validOrderTypes = map[string]struct{}{
		typeLimit:      {},
		typeMarket:     {},
		typeStopLimit:  {},
		typeStop:       {},
		typeTakeProfit: {},
	}

	validOrderSides = map[string]struct{}{
		sideBid: {},
		sideAsk: {},
	}

	cancellableStatusMap = map[string]string{
		statusAccepted:         statusCancelled,
		statusPlaced:           statusCancelled,
		statusPartiallyMatched: statusPartiallyCancelled,
	}

	openStatuses = []string{
		statusAccepted,
		statusPlaced,
		statusPartiallyMatched,
		statusFullyMatched,
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
