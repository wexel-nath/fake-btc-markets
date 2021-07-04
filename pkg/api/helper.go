package api

import (
	"net/http"
	"time"

	"fake-btc-markets/pkg/log"
)

func getTimestampFromRequest(r *http.Request) time.Time {
	timestampString := r.URL.Query().Get("timestamp")
	if timestampString != "" {
		t, err := time.Parse(time.RFC3339, timestampString)
		if err != nil {
			log.Error(err)
		} else {
			return t
		}
	}

	return time.Now()
}
