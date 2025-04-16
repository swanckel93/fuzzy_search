package searchCache

import (
	"sync"
	"github.com/swanckel93/fuzzy_api/search" // adjust import according to your project structure
)

// CacheKey is a tuple of Document ID and Query used as a key in the cache
type CacheKey struct {
	DocID string
	Query string
}

// cacheEntry holds the actual data in the cache
type cacheEntry struct {
	key    CacheKey
	value  []search.SearchResult
	size   int // in bytes
}

type SearchCache struct {
	mu          sync.Mutex
	data        map[CacheKey]*cacheEntry
	order       []CacheKey
	currentSize int
	maxSize     int // in bytes
}

// NewSearchCache creates a new search cache with a given max size in MB
func NewSearchCache(maxSizeMB int) *SearchCache {
	return &SearchCache{
		data:    make(map[CacheKey]*cacheEntry),
		order:   []CacheKey{},
		maxSize: maxSizeMB * 1024 * 1024,
	}
}

// Get retrieves search results from the cache by document ID and query
func (c *SearchCache) Get(docID, query string) ([]search.SearchResult, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	key := CacheKey{DocID: docID, Query: query}
	if entry, ok := c.data[key]; ok {
		// Move to the end (most recent)
		c.moveToEnd(key)
		return entry.value, true
	}
	return nil, false
}

// Set adds search results to the cache, evicting old entries if necessary
func (c *SearchCache) Set(docID, query string, results []search.SearchResult) {
	c.mu.Lock()
	defer c.mu.Unlock()

	key := CacheKey{DocID: docID, Query: query}

	size := estimateSize(results)
	if size > c.maxSize {
		return // entry too big to cache
	}

	// If already exists, remove and update
	if old, ok := c.data[key]; ok {
		c.remove(key, old.size)
	}

	// Evict if needed
	for c.currentSize+size > c.maxSize && len(c.order) > 0 {
		oldestKey := c.order[0]
		c.remove(oldestKey, c.data[oldestKey].size)
	}

	// Insert new entry
	c.data[key] = &cacheEntry{key: key, value: results, size: size}
	c.order = append(c.order, key)
	c.currentSize += size
}

// moveToEnd moves the given key to the end of the access order slice (most recent)
func (c *SearchCache) moveToEnd(key CacheKey) {
	for i, k := range c.order {
		if k == key {
			// Move to end
			c.order = append(c.order[:i], c.order[i+1:]...)
			c.order = append(c.order, key)
			return
		}
	}
}

// remove removes the key and updates the current size
func (c *SearchCache) remove(key CacheKey, size int) {
	delete(c.data, key)
	for i, k := range c.order {
		if k == key {
			c.order = append(c.order[:i], c.order[i+1:]...)
			break
		}
	}
	c.currentSize -= size
}

// estimateSize estimates the size of the search results in bytes
func estimateSize(results []search.SearchResult) int {
	size := 0
	for _, r := range results {
		// Approximate size: 16 bytes overhead + string lengths + ints
		size += 16 + len(r.Sentence) + len(r.Match) + 4*8 // 8 bytes for each int field
	}
	return size
}
