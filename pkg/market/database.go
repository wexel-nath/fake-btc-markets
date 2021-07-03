package market

import (
	"strings"
	"time"

	"fake-btc-markets/pkg/database"
)

const (
	// market columns
	columnMarketID       = "market_id"
	columnBaseAsset      = "base_asset"
	columnQuoteAsset     = "quote_asset"
	columnMinOrderAmount = "min_order_amount"
	columnMaxOrderAmount = "max_order_amount"
	columnAmountDecimals = "amount_decimals"
	columnPriceDecimals  = "price_decimals"

	// market period columns
	columnMarketPeriodID  = "market_period_id"
	columnTimePeriodStart = "time_period_start"
	columnTimePeriodEnd   = "time_period_end"
	columnPriceOpen       = "price_open"
	columnPriceHigh       = "price_high"
	columnPriceLow        = "price_low"
	columnPriceClose      = "price_close"
	columnVolumeTraded    = "volume_traded"

	// ticker columns
	columnBestBid    = "best_bid"
	columnBestAsk    = "best_ask"
	columnLastPrice  = "last_price"
	columnVolume24h  = "volume24h"
	columnPrice24h   = "price24h"
	columnLow24h     = "low24h"
	columnHigh24h    = "high24h"
	columnTimestamp  = "timestamp"
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

	marketPeriodColumns = []string{
		columnMarketPeriodID,
		columnMarketID,
		columnTimePeriodStart,
		columnTimePeriodEnd,
		columnPriceOpen,
		columnPriceHigh,
		columnPriceLow,
		columnPriceClose,
		columnVolumeTraded,
	}

	marketPeriodColumnsString = strings.Join(marketPeriodColumns, ", ")

	marketTickerColumns = []string{
		columnMarketID,
		columnBestBid,
		columnBestAsk,
		columnLastPrice,
		columnVolume24h,
		columnPrice24h,
		columnLow24h,
		columnHigh24h,
		columnTimestamp,
	}
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

func insertPeriod(
	marketID string,
	timePeriodStart string,
	timePeriodEnd string,
	priceOpen float64,
	priceHigh float64,
	priceLow float64,
	priceClose float64,
	volumeTraded float64,
) (map[string]interface{}, error) {
	query := `
		INSERT INTO market_period (` + marketPeriodColumnsString + `)
		VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING ` + marketPeriodColumnsString

	params := []interface{}{
		marketID,
		timePeriodStart,
		timePeriodEnd,
		priceOpen,
		priceHigh,
		priceLow,
		priceClose,
		volumeTraded,
	}

	return database.QueryRow(query, params, marketPeriodColumns)
}

func selectMarketTicker(marketID string, timestamp time.Time) (map[string]interface{}, error) {
	openingTimestamp := timestamp.Add(-24 * time.Hour)
	query := `
		WITH ticker_period AS (
			SELECT *
			FROM market_period
			WHERE market_id = $1
			AND time_period_end BETWEEN $2 AND $3
		),
		opening_period AS (
			SELECT *
			FROM ticker_period
			ORDER BY time_period_end
			LIMIT 1
		),
		closing_period AS (
			SELECT *
			FROM ticker_period
			ORDER BY time_period_end DESC
			LIMIT 1
		),
		period_aggregate AS (
			SELECT
				market_id,
				SUM(volume_traded) AS volume24h,
				MIN(price_low) AS low24h,
				MAX(price_high) AS high24h
			FROM ticker_period
			GROUP BY market_id
		)
		SELECT
			market_id,
			closing_period.price_close AS best_bid,
			closing_period.price_close AS best_ask,
			closing_period.price_close AS last_price,
			period_aggregate.volume24h,
			closing_period.price_close - opening_period.price_close AS price24h,
			period_aggregate.low24h,
			period_aggregate.high24h,
			closing_period.time_period_end AS timestamp
		FROM opening_period
			JOIN closing_period USING (market_id)
			JOIN period_aggregate USING (market_id)
	`

	params := []interface{}{
		marketID,
		openingTimestamp,
		timestamp,
	}

	return database.QueryRow(query, params, marketTickerColumns)
}

func selectLatestPeriodForMarket(marketID string) (map[string]interface{}, error) {
	query := `
		SELECT ` + marketPeriodColumnsString + `
		FROM market_period
		WHERE market_id = $1
		ORDER BY time_period_end DESC
		LIMIT 1
	`

	params := []interface{}{
		marketID,
	}

	return database.QueryRow(query, params, marketPeriodColumns)
}
