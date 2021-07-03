package order

import (
	"strings"

	"fake-btc-markets/pkg/database"
)

const (
	columnMarketID           = "market_id"
	columnOrderID            = "order_id"
	columnOrderPrice         = "order_price"
	columnOrderAmount        = "order_amount"
	columnOrderType          = "order_type"
	columnOrderSide          = "order_side"
	columnOrderTriggerPrice  = "order_trigger_price"
	columnOrderTriggerAmount = "order_trigger_amount"
	columnOrderTimeInForce   = "order_time_in_force"
	columnOrderPostOnly      = "order_post_only"
	columnOrderSelfTrade     = "order_self_trade"
	columnOrderCreated       = "order_created"
	columnClientOrderID      = "client_order_id"
)

var (
	orderColumns = []string{
		columnMarketID,
		columnOrderID,
		columnOrderPrice,
		columnOrderAmount,
		columnOrderType,
		columnOrderSide,
		columnOrderTriggerPrice,
		columnOrderTriggerAmount,
		columnOrderTimeInForce,
		columnOrderPostOnly,
		columnOrderSelfTrade,
		columnOrderCreated,
		columnClientOrderID,
	}

	orderColumnsString = strings.Join(orderColumns, ", ")
)

func insertSimpleOrder(
	marketID string,
	price float64,
	amount float64,
	orderType string,
	side string,
) (map[string]interface{}, error) {
	query := `
		INSERT INTO "order" (` + orderColumnsString + `)
		VALUES (
			$1,
			DEFAULT,
			$2,
			$3,
			$4,
			$5,
			NULL,
			NULL,
			DEFAULT,
			DEFAULT,
			DEFAULT,
			DEFAULT,
			NULL
		)
		RETURNING ` + orderColumnsString

	params := []interface{}{
		marketID,
		price,
		amount,
		orderType,
		side,
	}

	return database.QueryRow(query, params, orderColumns)
}
