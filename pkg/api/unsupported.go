package api

import (
	"net/http"

	"fake-btc-markets/pkg/api/request"
)

func unsupported(_ http.ResponseWriter, _ *http.Request) request.Response {
	return request.Response{
		Data: nil,
		Meta: &request.Meta{
			Message: "This endpoint has not been implemented yet",
		},
		StatusCode: http.StatusBadRequest,
	}
}
