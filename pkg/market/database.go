package market

import (
	"strings"

	"fake-btc-markets/pkg/database"
)

const (
	columnMarketID       = "market_id"
	columnBaseAsset      = "base_asset"
	columnQuoteAsset     = "quote_asset"
	columnMinOrderAmount = "min_order_amount"
	columnMaxOrderAmount = "max_order_amount"
	columnAmountDecimals = "amount_decimals"
	columnPriceDecimals  = "price_decimals"
)

var (
	marketColumns = []string{
		columnMarketID,
		columnBaseAsset,
		columnQuoteAsset,
		columnMinOrderAmount,
		columnMaxOrderAmount,
		columnAmountDecimals,
		columnPriceDecimals,
	}

	marketColumnsString = strings.Join(marketColumns, ", ")
)

func insertMarket(
	id string,
	baseAsset string,
	quoteAsset string,
	minOrderAmount string,
	maxOrderAmount string,
	amountDecimals string,
	priceDecimals string,
) (map[string]interface{}, error) {
	query := `
		INSERT INTO market (` + marketColumnsString + `)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING ` + marketColumnsString

	params := []interface{}{
		id,
		baseAsset,
		quoteAsset,
		minOrderAmount,
		maxOrderAmount,
		amountDecimals,
		priceDecimals,
	}

	return database.QueryRow(query, params, marketColumns)
}

func selectMarket(id string) (map[string]interface{}, error) {
	query := `
		SELECT ` + marketColumnsString + `
		FROM market
		WHERE ` + columnMarketID + ` = $1
	`
	params := []interface{}{
		id,
	}

	return database.QueryRow(query, params, marketColumns)
}

func selectMarkets() ([]map[string]interface{}, error) {
	query := `
		SELECT ` + marketColumnsString + `
		FROM market
	`

	return database.QueryRows(query, nil, marketColumns)
}
