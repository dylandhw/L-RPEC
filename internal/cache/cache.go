package cache

import (
	"net/http"
	"time"
)

type Entry struct {
	ResponseBody []byte
	http.Header
	StatusCode int
	ExpiryTime time.Time
}
