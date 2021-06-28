package coinapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"fake-btc-markets/pkg/config"
)

const (
	baseURL = "https://rest.coinapi.io"
	limit   = 10000
)

type Period struct{
	Start  string  `json:"time_period_start"`
	End    string  `json:"time_period_end"`
	Open   float64 `json:"rate_open"`
	High   float64 `json:"rate_high"`
	Low    float64 `json:"rate_low"`
	Close  float64 `json:"rate_close"`
	Volume float64 `json:"volume_traded"`
}

func GetHistoricalData(baseAsset string, quoteAsset string, timeStart time.Time) ([]Period, error) {
	url := fmt.Sprintf(
		"%s/v1/exchangerate/%s/%s/history?period_id=10MIN&time_start=%s&limit=%d",
		baseURL,
		baseAsset,
		quoteAsset,
		timeStart.Format("2006-01-02T15:04:05"),
		limit,
	)

	response, err := callCoinApi(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s returned %d, expected 200", url, response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	var periods []Period
	err = json.Unmarshal(body, &periods)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling %v returned err: %v", string(body), err)
	}

	return periods, nil
}

func callCoinApi(url string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("X-CoinAPI-Key", config.Get().CoinApiKey)

	client := http.Client{
		Timeout: 30 * time.Second,
	}
	return client.Do(request)
}
