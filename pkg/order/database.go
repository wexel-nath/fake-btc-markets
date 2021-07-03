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
	columnOrderStatus        = "order_status"
	columnClientOrderID      = "client_order_id"
	columnOpenAmount         = "open_amount"
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
		columnOrderStatus,
		columnClientOrderID,
		columnOpenAmount,
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
		WITH insert_order AS (
			INSERT INTO "order" (
				market_id,
				order_price,
				order_amount,
				order_type,
				order_side
			)
			VALUES (
				$1,
				$2,
				$3,
				$4,
				$5
			)
			RETURNING *
		),
		order_open_amount AS (
			SELECT
				order_id,
				order_amount AS open_amount
			FROM insert_order
		)
		SELECT ` + orderColumnsString + `
		FROM insert_order
		JOIN order_open_amount USING(order_id)
	`

	params := []interface{}{
		marketID,
		price,
		amount,
		orderType,
		side,
	}

	return database.QueryRow(query, params, orderColumns)
}

func selectOrder(orderID int64) (map[string]interface{}, error) {
	query := `
		WITH order_open_amount AS (
			SELECT
				order_id,
				order_amount - COALESCE(SUM(trade_amount), 0) AS open_amount
			FROM "order"
			LEFT JOIN trade USING(order_id)
			WHERE order_id = $1
			GROUP BY order_id, order_amount
		)
		SELECT ` + orderColumnsString + `
		FROM "order"
		JOIN order_open_amount USING(order_id)
		WHERE order_id = $1
	`

	params := []interface{}{
		orderID,
	}

	return database.QueryRow(query, params, orderColumns)
}

func selectOrders() ([]map[string]interface{}, error) {
	query := `
		WITH order_open_amount AS (
			SELECT
				order_id,
				order_amount - COALESCE(SUM(trade_amount), 0) AS open_amount
			FROM "order"
			LEFT JOIN trade USING(order_id)
			GROUP BY order_id, order_amount
		)
		SELECT ` + orderColumnsString + `
		FROM "order"
		JOIN order_open_amount USING(order_id)
		ORDER BY order_created DESC
	`

	return database.QueryRows(query, nil, orderColumns)
}

func updateOrderStatus(orderID int64, status string) (map[string]interface{}, error) {
	columns := []string{
		columnOrderID,
		columnClientOrderID,
	}

	query := `
		UPDATE "order"
		SET order_status = $2
		WHERE order_id = $1
		RETURNING ` + strings.Join(columns, ", ")

	params := []interface{}{
		orderID,
		status,
	}

	return database.QueryRow(query, params, columns)
}
