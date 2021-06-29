package market

import (
	"fmt"
	"time"
)

type Ticker struct{
	MarketID  string    `json:"marketId"`
	BestBid   string    `json:"bestBid"`
	BestAsk   string    `json:"bestAsk"`
	LastPrice string    `json:"lastPrice"`
	Volume24h string    `json:"volume24h"`
	Price24h  string    `json:"price24h"`
	Low24h    string    `json:"low24h"`
	High24h   string    `json:"high24h"`
	Timestamp time.Time `json:"timestamp"`
}

func newTickerFromRow(row map[string]interface{}) (t Ticker, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic[%v] while building ticker from row[%v]", r, row)
		}
	}()

	ticker := Ticker{
		MarketID:  row[columnMarketID].(string),
		BestBid:   string(row["best_bid"].([]byte)),
		BestAsk:   string(row["best_ask"].([]byte)),
		LastPrice: string(row["last_price"].([]byte)),
		Volume24h: string(row["volume24h"].([]byte)),
		Price24h:  string(row["price24h"].([]byte)),
		Low24h:    string(row["low24h"].([]byte)),
		High24h:   string(row["high24h"].([]byte)),
		Timestamp: row["timestamp"].(time.Time),
	}
	return ticker, nil
}

func GetTickerForTimestamp(marketID string, timestamp time.Time) (Ticker, error) {
	row, err := selectMarketTicker(marketID, timestamp)
	if err != nil {
		return Ticker{}, err
	}

	return newTickerFromRow(row)
}

func GetLatestTicker(marketID string) (Ticker, error) {
	period, err := GetLatestPeriodForMarket(marketID)
	if err != nil {
		return Ticker{}, err
	}

	return GetTickerForTimestamp(marketID, period.TimePeriodEnd)
}
