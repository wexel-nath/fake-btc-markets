package market

import (
	"fmt"
	"strconv"
	"time"

	"fake-btc-markets/pkg/helper/parse"
	"fake-btc-markets/pkg/log"
)

type Period struct{
	ID              int64
	MarketID        string
	TimePeriodStart time.Time
	TimePeriodEnd   time.Time
	PriceOpen       float64
	PriceHigh       float64
	PriceLow        float64
	PriceClose      float64
	VolumeTraded    float64
}

func newPeriodFromRow(row map[string]interface{}) (p Period, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic[%v] while building period from row[%v]", r, row)
		}
	}()

	period := Period{
		ID:              row[columnMarketPeriodID].(int64),
		MarketID:        row[columnMarketID].(string),
		TimePeriodStart: row[columnTimePeriodStart].(time.Time),
		TimePeriodEnd:   row[columnTimePeriodEnd].(time.Time),
		PriceOpen:       mustParseFloat(row[columnPriceOpen]),
		PriceHigh:       mustParseFloat(row[columnPriceHigh]),
		PriceLow:        mustParseFloat(row[columnPriceLow]),
		PriceClose:      mustParseFloat(row[columnPriceClose]),
		VolumeTraded:    mustParseFloat(row[columnVolumeTraded]),
	}

	return period, nil
}

func mustParseFloat(val interface{}) float64 {
	floatVal, err := strconv.ParseFloat(parse.BytesAsString(val), 64)
	if err != nil {
		panic(err)
	}

	return floatVal
}

func newPeriodsFromRows(rows []map[string]interface{}) []Period {
	periods := make([]Period, 0)
	for _, row := range rows {
		period, err := newPeriodFromRow(row)
		if err != nil {
			log.Error(err)
			continue
		}

		periods = append(periods, period)
	}

	return periods
}

func NewPeriod(
	marketID string,
	timePeriodStart string,
	timePeriodEnd string,
	priceOpen float64,
	priceHigh float64,
	priceLow float64,
	priceClose float64,
	volumeTraded float64,
) (Period, error) {
	row, err := insertPeriod(
		marketID,
		timePeriodStart,
		timePeriodEnd,
		priceOpen,
		priceHigh,
		priceLow,
		priceClose,
		volumeTraded,
	)
	if err != nil {
		return Period{}, err
	}

	return newPeriodFromRow(row)
}

func GetLatestPeriodForMarket(marketID string) (Period, error) {
	row, err := selectLatestPeriodForMarket(marketID)
	if err != nil {
		return Period{}, err
	}

	return newPeriodFromRow(row)
}
