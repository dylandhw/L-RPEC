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
	mutex   sync.Mutex
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

func (cache *Cache) Get(key string) (Entry, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	entry, ok := cache.Entries[key]
	return entry, ok
}

func (cache *Cache) Set(key string, entry Entry) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.Entries[key] = entry
}
