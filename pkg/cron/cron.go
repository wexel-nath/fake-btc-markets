package cron

import (
	"database/sql"
	"time"

	"fake-btc-markets/pkg/coinapi"
	"fake-btc-markets/pkg/log"
	"fake-btc-markets/pkg/market"
)

const (
	defaultInitialDate = "2016-01-01T00:00:00Z"

	baseAsset  = "ETH"
	quoteAsset = "USD"
)

func DoGetHistoricalData() {
	marketID := baseAsset + "-" + quoteAsset
	log.Info("Getting historical data for %s", marketID)

	err := getHistoricalData(marketID)
	if err != nil {
		log.Error(err)
	}

	log.Info("Finished getting historical data for %s", marketID)
}

func getHistoricalData(marketID string) error {
	timeStart, err := time.Parse(time.RFC3339, defaultInitialDate)
	if err != nil {
		return err
	}

	period, err := market.GetLatestPeriodForMarket(marketID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == nil {
		// start from end of the latest period
		timeStart = period.TimePeriodEnd
	}

	periods, err := coinapi.GetHistoricalData(baseAsset, quoteAsset, timeStart)
	if err != nil {
		return err
	}

	// insert periods
	for _, p := range periods {
		_, err = market.NewPeriod(
			marketID,
			p.Start,
			p.End,
			p.Open,
			p.High,
			p.Low,
			p.Close,
			p.Volume,
		)
		if err != nil {
			// be defensive and exit, for now
			return err
		}
	}

	return nil
}
