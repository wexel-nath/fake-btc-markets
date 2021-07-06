package order

import (
	"fmt"
	"strings"
	"time"

	"fake-btc-markets/pkg/helper/parse"
	"fake-btc-markets/pkg/log"
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
	Status        string    `json:"status"`
	OpenAmount    string    `json:"openAmount"`
}

func newOrderFromRow(row map[string]interface{}) (o Order, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic[%v] while building order from row[%v]", r, row)
		}
	}()

	order := Order{
		ID:           parse.IntAsString(row[columnOrderID]),
		MarketID:     row[columnMarketID].(string),
		Price:        parse.BytesAsString(row[columnOrderPrice]),
		Amount:       parse.BytesAsString(row[columnOrderAmount]),
		Side:         row[columnOrderSide].(string),
		Type:         row[columnOrderType].(string),
		TimeInForce:  row[columnOrderTimeInForce].(string),
		PostOnly:     row[columnOrderPostOnly].(bool),
		SelfTrade:    row[columnOrderSelfTrade].(string),
		CreationTime: row[columnOrderCreated].(time.Time),
		Status:       row[columnOrderStatus].(string),
		OpenAmount:   parse.BytesAsString(row[columnOpenAmount]),
	}

	if !order.isOpen() {
		order.OpenAmount = "0"
	}

	// nullable fields

	triggerPrice := row[columnOrderTriggerPrice]
	if triggerPrice != nil {
		*order.TriggerPrice = parse.BytesAsString(triggerPrice)
	}

	triggerAmount := row[columnOrderTriggerAmount]
	if triggerAmount != nil {
		*order.TriggerAmount = parse.BytesAsString(triggerAmount)
	}

	clientOrderID := row[columnClientOrderID]
	if clientOrderID != nil {
		*order.ClientOrderID = clientOrderID.(string)
	}

	return order, nil
}

func (o Order) isOpen() bool {
	for _, s := range openStatuses {
		if s == o.Status {
			return true
		}
	}
	return false
}

func (o Order) validateRequest() []string {
	errors := make([]string, 0)

	_, err := market.GetMarketByID(o.MarketID)
	if err != nil {
		errors = append(errors, "invalid marketId")
	}

	price, err := parse.StringToFloat(o.Price)
	if err != nil || price <= 0 {
		errors = append(errors, "invalid price")
	}

	amount, err := parse.StringToFloat(o.Amount)
	if err != nil || amount <= 0 {
		errors = append(errors, "invalid amount")
	}

	if ok := isValidOrderType(o.Type); !ok {
		errors = append(errors, "invalid orderType")
	}

	if ok := isValidOrderSide(o.Side); !ok {
		errors = append(errors, "invalid orderSide")
	}

	return errors
}

func NewOrder(o Order, timestamp time.Time) (Order, error) {
	errors := o.validateRequest()
	if len(errors) > 0 {
		return Order{}, fmt.Errorf(strings.Join(errors, ". "))
	}

	row, err := insertSimpleOrder(
		o.MarketID,
		parse.MustGetFloat(o.Price),
		parse.MustGetFloat(o.Amount),
		o.Type,
		o.Side,
		timestamp,
	)
	if err != nil {
		return Order{}, err
	}

	order, err := newOrderFromRow(row)
	if err != nil {
		return Order{}, err
	}

	trade, err := maybeCreateTrade(order)
	if err != nil {
		log.Error(err)
		return order, nil
	}

	if parse.MustGetInt(trade.ID) > 0 {
		// trade was executed, re-fetch order
		return GetOrderByID(order.ID)
	}

	return order, nil
}

func GetOrderByID(orderID string) (Order, error) {
	id, err := parse.StringToInt(orderID)
	if err != nil {
		return Order{}, err
	}

	row, err := selectOrder(id)
	if err != nil {
		return Order{}, err
	}

	return newOrderFromRow(row)
}

func GetOrders() ([]Order, error) {
	rows, err := selectOrders()
	if err != nil {
		return nil, err
	}

	orders := make([]Order, 0)
	for _, row := range rows {
		order, err := newOrderFromRow(row)
		if err != nil {
			log.Error(err)
			continue
		}

		orders = append(orders, order)
	}

	return orders, nil
}

type CancelledOrder struct{
	OrderID       string  `json:"orderId"`
	ClientOrderID *string `json:"clientOrderId"`
}

func newCancelledOrderFromRow(row map[string]interface{}) (c CancelledOrder, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic[%v] while building cancelled order from row[%v]", r, row)
		}
	}()

	c = CancelledOrder{
		OrderID: parse.IntAsString(row[columnOrderID]),
	}

	clientOrderID := row[columnClientOrderID]
	if clientOrderID != nil {
		*c.ClientOrderID = clientOrderID.(string)
	}

	return c, nil
}

func CancelOrder(orderID string) (CancelledOrder, error) {
	order, err := GetOrderByID(orderID)
	if err != nil {
		return CancelledOrder{}, err
	}

	cancelStatus, ok := cancellableStatusMap[order.Status]
	if !ok {
		return CancelledOrder{}, fmt.Errorf("order status %s cannot be cancelled", order.Status)
	}

	id, err := parse.StringToInt(orderID)
	if err != nil {
		return CancelledOrder{}, err
	}

	row, err := updateOrderStatus(id, cancelStatus)
	if err != nil {
		return CancelledOrder{}, err
	}

	return newCancelledOrderFromRow(row)
}
