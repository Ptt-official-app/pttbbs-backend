package utils

import (
	"net/http"
)

// httpClient has no Client.Timeout on purpose: the per-request timeout is
// types.EXPIRE_HTTP_REQUEST_TS, applied through the request context in httpProcess.
var httpClient = &http.Client{}
