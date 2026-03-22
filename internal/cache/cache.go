package cache

import (
	"net/http"
	"sync"
	"time"
)

type Entry struct {
	ResponseBody []byte
	Headers      http.Header
	StatusCode   int
	ExpiryTime   time.Time
}

type Cache struct {
	mu      sync.Mutex
	Entries map[string]Entry
}

func NewEntry(ResponseBody []byte, Headers http.Header, StatusCode int, ExpiryTime time.Time) *Entry {
	return &Entry{
		ResponseBody: ResponseBody,
		Headers:      Headers,
		StatusCode:   StatusCode,
		ExpiryTime:   ExpiryTime,
	}
}
