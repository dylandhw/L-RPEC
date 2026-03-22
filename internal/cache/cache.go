package cache

import (
	"net/http"
	"sync"
	"time"
)

type Entry struct {
	ResponseBody []byte
	http.Header
	StatusCode int
	ExpiryTime time.Time
}

type Cache struct {
	mu      sync.Mutex
	Entries map[string]Entry
}
