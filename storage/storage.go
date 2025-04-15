package storage

import (
	"sync"
)

type DocumentStore struct {
	mu    sync.RWMutex
	Files map[string][]string
}

var store = &DocumentStore{
	Files: make(map[string][]string),
}

func AddFile(filename string, sentences []string) {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.Files[filename] = sentences
}

func GetFile(filename string) ([]string, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	sentences, ok := store.Files[filename]
	return sentences, ok
}

func ListFiles() []string {
	store.mu.RLock()
	defer store.mu.RUnlock()
	keys := make([]string, 0, len(store.Files))
	for k := range store.Files {
		keys = append(keys, k)
	}
	return keys
}
