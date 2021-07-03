package order

import (
	"fmt"
	"strings"
	"time"

	"fake-btc-markets/pkg/helper"
	"fake-btc-markets/pkg/market"
)

type Order struct{
	ID            string    `json:"orderId"`
	MarketID      string    `json:"marketId"`
	Price         string    `json:"price"`
	Amount        string    `json:"amount"`
	Side          string    `json:"side"`
	Type          string    `json:"type"`
	TriggerPrice  *string   `json:"triggerPrice"`
	TriggerAmount *string   `json:"triggerAmount"`
	TimeInForce   string    `json:"timeInForce"`
	PostOnly      bool      `json:"postOnly"`
	SelfTrade     string    `json:"selfTrade"`
	ClientOrderID *string   `json:"clientOrderId"`
	CreationTime  time.Time `json:"creationTime"`

	// derived from trades
	OpenAmount string `json:"openAmount"`
	Status     string `json:"status"`
}

func newOrderFromRow(row map[string]interface{}) (o Order, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic[%v] while building order from row[%v]", r, row)
		}
	}()

	order := Order{
		ID:             helper.ParseIntAsString(row[columnOrderID]),
		MarketID:       row[columnMarketID].(string),
		Price:          string(row[columnOrderPrice].([]byte)),
		Amount:         string(row[columnOrderAmount].([]byte)),
		Side:           row[columnOrderSide].(string),
		Type:           row[columnOrderType].(string),
		TimeInForce:    row[columnOrderTimeInForce].(string),
		PostOnly:       row[columnOrderPostOnly].(bool),
		SelfTrade:      row[columnOrderSelfTrade].(string),
		CreationTime:   row[columnOrderCreated].(time.Time),

		// these should be derived from trades
		OpenAmount: string(row[columnOrderAmount].([]byte)),
		Status:     statusAccepted,
	}

	// nullable fields

	triggerPrice := row[columnOrderTriggerPrice]
	if triggerPrice != nil {
		*order.TriggerPrice = string(triggerPrice.([]byte))
	}

	triggerAmount := row[columnOrderTriggerAmount]
	if triggerAmount != nil {
		*order.TriggerAmount = string(triggerAmount.([]byte))
	}

	clientOrderID := row[columnClientOrderID]
	if clientOrderID != nil {
		*order.ClientOrderID = clientOrderID.(string)
	}

	return order, nil
}

func (o Order) validateRequest() []string {
	errors := make([]string, 0)

	_, err := market.GetMarketByID(o.MarketID)
	if err != nil {
		errors = append(errors, fmt.Sprintf("invalid marketId. err: %v", err))
	}

	_, err = helper.StringToFloat(o.Price)
	if err != nil {
		errors = append(errors, fmt.Sprintf("invalid price. err: %v", err))
	}

	_, err = helper.StringToFloat(o.Amount)
	if err != nil {
		errors = append(errors, fmt.Sprintf("invalid amount. err: %v", err))
	}

	if ok := isValidOrderType(o.Type); !ok {
		errors = append(errors, "invalid orderType")
	}

	if ok := isValidOrderSide(o.Side); !ok {
		errors = append(errors, "invalid orderSide")
	}

	return errors
}

func NewOrder(o Order) (Order, error) {
	errors := o.validateRequest()
	if len(errors) > 0 {
		return Order{}, fmt.Errorf(strings.Join(errors, ". "))
	}

	row, err := insertSimpleOrder(
		o.MarketID,
		helper.MustGetFloat(o.Price),
		helper.MustGetFloat(o.Amount),
		o.Type,
		o.Side,
	)
	if err != nil {
		return Order{}, err
	}

	// todo: find/place trades for order

	return newOrderFromRow(row)
}
