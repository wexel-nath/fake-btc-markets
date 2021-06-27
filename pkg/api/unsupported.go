package api

import (
	"net/http"
)

func unsupported(_ http.ResponseWriter, _ *http.Request) response {
	meta := newMeta("This endpoint has not been implemented yet")
	return newResponse(nil, meta, http.StatusBadRequest)
}
