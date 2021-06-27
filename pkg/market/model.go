package market

import (
	"fake-btc-markets/pkg/log"

	"github.com/mitchellh/mapstructure"
)

type Market struct {
	ID             string `json:"market_id" mapstructure:"market_id"`
	BaseAsset      string `json:"base_asset" mapstructure:"base_asset"`
	QuoteAsset     string `json:"quote_asset" mapstructure:"quote_asset"`
	MinOrderAmount string `json:"min_order_amount" mapstructure:"min_order_amount"`
	MaxOrderAmount string `json:"max_order_amount" mapstructure:"max_order_amount"`
	AmountDecimals string `json:"amount_decimals" mapstructure:"amount_decimals"`
	PriceDecimals  string `json:"price_decimals" mapstructure:"price_decimals"`
}

func newMarketFromRow(row map[string]interface{}) (Market, error) {
	market := Market{}
	err := mapstructure.Decode(row, &market)
	return market, err
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
