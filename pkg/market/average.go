package market

import (
	"fmt"
	"regexp"
	"time"

	"fake-btc-markets/pkg/helper/parse"
	"fake-btc-markets/pkg/log"
)

var (
	// moving average formats
	maRegex = regexp.MustCompile("(sma).([0-9]+)([hd])")

	intervalMap = map[string]string{
		"d": "days",
		"h": "hours",
	}
)

type movingAverage struct{
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func newMovingAverage(name string, value float64) movingAverage {
	return movingAverage{
		Name:  name,
		Value: value,
	}
}

func getMovingAverages(marketID string, timestamp time.Time, maNames []string) []movingAverage {
	mas := make([]movingAverage, 0)

	for _, name := range maNames {
		ma, err := getMovingAverageForTimestamp(marketID, timestamp, name)
		if err != nil {
			log.Error(fmt.Errorf("error getting moving average[%s] for market[%s]: %v", name, marketID, err))
			continue
		}

		mas = append(mas, ma)
	}

	return mas
}

func getMovingAverageForTimestamp(marketID string, timestamp time.Time, name string) (movingAverage, error) {
	// todo: support averages other than SMA's

	_, interval, err := getTypeAndIntervalFromName(name)
	if err != nil {
		return movingAverage{}, err
	}

	row, err := selectMarketSimpleMovingAverage(marketID, timestamp, interval)
	if err != nil {
		return movingAverage{}, err
	}

	return newMovingAverage(name, parse.MustParseFloat(row[columnMovingAverage])), nil
}

func getTypeAndIntervalFromName(name string) (string, string, error) {
	matches := maRegex.FindStringSubmatch(name)
	if len(matches) != 4 {
		return "", "", fmt.Errorf("moving average[%s] expected 4 submatches, got %d", name, len(matches))
	}

	maType := matches[1]
	durationAmount := matches[2]
	durationType := matches[3]

	duration, ok := intervalMap[durationType]
	if !ok {
		return "", "", fmt.Errorf("moving average[%s] duration type[%s] is invalid", name, durationType)
	}

	interval := durationAmount + " " + duration
	return maType, interval, nil
}
