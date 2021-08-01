package order

import (
	"fmt"
	"time"

	"fake-btc-markets/pkg/fee"
	"fake-btc-markets/pkg/helper/parse"
	"fake-btc-markets/pkg/log"
	"fake-btc-markets/pkg/market"
)

type Trade struct{
	ID            string    `json:"id"`
	MarketID      string    `json:"marketId"`
	Timestamp     time.Time `json:"timestamp"`
	Price         string    `json:"price"`
	Amount        string    `json:"amount"`
	Side          string    `json:"side"`
	Fee           string    `json:"fee"`
	OrderID       string    `json:"orderId"`
	LiquidityType string    `json:"liquidityType"`
	ClientOrderID *string   `json:"clientOrderId"`
}

func newTradeFromRow(row map[string]interface{}) (t Trade, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic[%v] while building trade from row[%v]", r, row)
		}
	}()

	trade := Trade{
		ID:            parse.IntAsString(row[columnTradeID]),
		MarketID:      row[columnMarketID].(string),
		Timestamp:     row[columnTradeCreated].(time.Time),
		Price:         parse.BytesAsString(row[columnOrderPrice]),
		Amount:        parse.BytesAsString(row[columnTradeAmount]),
		Side:          row[columnOrderSide].(string),
		Fee:           parse.BytesAsString(row[columnTradeFee]),
		OrderID:       parse.IntAsString(row[columnOrderID]),
		LiquidityType: row[columnTradeLiquidityType].(string),
	}

	// nullable fields

	clientOrderID := row[columnClientOrderID]
	if clientOrderID != nil {
		*trade.ClientOrderID = clientOrderID.(string)
	}

	return trade, nil
}

func newTradesFromRows(rows []map[string]interface{}) ([]Trade, error) {
	trades := make([]Trade, 0)

	for _, row := range rows {
		trade, err := newTradeFromRow(row)
		if err != nil {
			log.Error(err)
			continue
		}

		trades = append(trades, trade)
	}

	return trades, nil
}

func NewTradeForOrder(order Order, timestamp time.Time) (Trade, error) {
	liquidityType := tradeLiquidityTypeTaker // constant for now
	amount := parse.MustGetFloat(order.Amount) // trade for the full amount
	traderVolume := 0.0 // todo: fetch from db
	orderID := parse.MustGetInt(order.ID)

	cost := parse.MustGetFloat(order.Price) * amount
	tradeFee := fee.CalculateTradeFee(cost, traderVolume)

	orderStatus := getMatchedStatus(amount, parse.MustGetFloat(order.OpenAmount))

	row, err := insertTrade(
		orderID,
		amount,
		tradeFee,
		liquidityType,
		timestamp,
		orderStatus,
	)
	if err != nil {
		return Trade{}, err
	}

	trade, err := newTradeFromRow(row)
	if err != nil {
		return Trade{}, err
	}

	return trade, nil
}

func getMatchedStatus(amount float64, openAmount float64) string {
	if amount == openAmount {
		return statusFullyMatched
	}
	return  statusPartiallyMatched
}

// maybeCreateTrade creates a trade if the order price matches the last price at the given timestamp
func maybeCreateTrade(order Order) (Trade, error) {
	ticker, err := market.GetTickerForTimestamp(order.MarketID, order.CreationTime, nil)
	if err != nil {
		return Trade{}, err
	}

	orderPrice := fmt.Sprintf("%.5f", parse.MustGetFloat(order.Price))
	lastPrice := fmt.Sprintf("%.5f", parse.MustGetFloat(ticker.LastPrice))

	// ticker should be within 10 minutes of order timestamp
	// and order price must match ticker last price
	if order.CreationTime.Sub(ticker.Timestamp) > 10 * time.Minute ||
		orderPrice != lastPrice {
		return Trade{}, nil
	}

	return NewTradeForOrder(order, order.CreationTime)
}

func GetTradesByOrderID(orderID string) ([]Trade, error) {
	id, err := parse.StringToInt(orderID)
	if err != nil {
		return nil, err
	}

	rows, err := selectTradesByOrderID(id)
	if err != nil {
		return nil, err
	}

	return newTradesFromRows(rows)
}
