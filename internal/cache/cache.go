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

func NewCache() *Cache {
	return &Cache{
		Entries: make(map[string]Entry),
	}
}

func Get(key string) (Entry, bool) {}

func Set(key string, entry Entry) {}
