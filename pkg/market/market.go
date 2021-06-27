package market

import (
	"fmt"

	"fake-btc-markets/pkg/log"
)

type Market struct{
	ID             string `json:"marketId"`
	BaseAsset      string `json:"baseAsset"`
	QuoteAsset     string `json:"quoteAsset"`
	MinOrderAmount string `json:"minOrderAmount"`
	MaxOrderAmount string `json:"maxOrderAmount"`
	AmountDecimals string `json:"amountDecimals"`
	PriceDecimals  string `json:"priceDecimals"`
}

func newMarketFromRow(row map[string]interface{}) (m Market, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic[%v] while building market from row[%v]", r, row)
		}
	}()

	market := Market{
		ID:             row[columnMarketID].(string),
		BaseAsset:      row[columnBaseAsset].(string),
		QuoteAsset:     row[columnQuoteAsset].(string),
		MinOrderAmount: row[columnMinOrderAmount].(string),
		MaxOrderAmount: row[columnMaxOrderAmount].(string),
		AmountDecimals: row[columnAmountDecimals].(string),
		PriceDecimals:  row[columnPriceDecimals].(string),
	}
	return market, nil
}

func newMarketsFromRows(rows []map[string]interface{}) []Market {
	markets := make([]Market, 0)
	for _, row := range rows {
		market, err := newMarketFromRow(row)
		if err != nil {
			log.Error(err)
			continue
		}

		markets = append(markets, market)
	}

	return markets
}

func NewMarket(
	baseAsset string,
	quoteAsset string,
	minOrderAmount string,
	maxOrderAmount string,
	amountDecimals string,
	priceDecimals string,
) (Market, error) {
	id := baseAsset + "-" + quoteAsset

	row, err := insertMarket(
		id,
		baseAsset,
		quoteAsset,
		minOrderAmount,
		maxOrderAmount,
		amountDecimals,
		priceDecimals,
	)
	if err != nil {
		return Market{}, err
	}

	return newMarketFromRow(row)
}

func GetMarketByID(id string) (Market, error) {
	row, err := selectMarket(id)
	if err != nil {
		return Market{}, err
	}

	return newMarketFromRow(row)
}

func GetMarkets() ([]Market, error) {
	rows, err := selectMarkets()
	if err != nil {
		return nil, err
	}

	return newMarketsFromRows(rows), nil
}
