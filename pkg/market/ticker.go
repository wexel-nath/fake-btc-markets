package market

import (
	"fmt"
	"time"

	"fake-btc-markets/pkg/helper/parse"
)

type Ticker struct{
	MarketID       string          `json:"marketId"`
	BestBid        string          `json:"bestBid"`
	BestAsk        string          `json:"bestAsk"`
	LastPrice      string          `json:"lastPrice"`
	Volume24h      string          `json:"volume24h"`
	Price24h       string          `json:"price24h"`
	Low24h         string          `json:"low24h"`
	High24h        string          `json:"high24h"`
	Timestamp      time.Time       `json:"timestamp"`
	MovingAverages []movingAverage `json:"moving_averages"`
}

func newTickerFromRow(row map[string]interface{}) (t Ticker, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic[%v] while building ticker from row[%v]", r, row)
		}
	}()

	ticker := Ticker{
		MarketID:  row[columnMarketID].(string),
		BestBid:   parse.BytesAsString(row[columnBestBid]),
		BestAsk:   parse.BytesAsString(row[columnBestAsk]),
		LastPrice: parse.BytesAsString(row[columnLastPrice]),
		Volume24h: parse.BytesAsString(row[columnVolume24h]),
		Price24h:  parse.BytesAsString(row[columnPrice24h]),
		Low24h:    parse.BytesAsString(row[columnLow24h]),
		High24h:   parse.BytesAsString(row[columnHigh24h]),
		Timestamp: row[columnTimestamp].(time.Time),
	}
	return ticker, nil
}

func GetTickerForTimestamp(marketID string, timestamp time.Time, maNames []string) (Ticker, error) {
	row, err := selectMarketTicker(marketID, timestamp)
	if err != nil {
		return Ticker{}, err
	}

	ticker, err := newTickerFromRow(row)
	if err != nil {
		return Ticker{}, err
	}

	ticker.MovingAverages = getMovingAverages(marketID, timestamp, maNames)
	return ticker, nil
}
