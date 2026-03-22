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

func (c *Cache) Get(key string) (Entry, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.Entries[key]
	return entry, ok
}

func (c *Cache) Set(key string, entry Entry) {
	c.mu.Lock()
	defer c.mu.Lock()
	c.Entries[key] = entry
}
