package main

import "sync"

var (
    store = make(map[string]string)
    mu    sync.RWMutex
)

func Save(shortID, longURL string) {
    mu.Lock()
    defer mu.Unlock()
    store[shortID] = longURL
}

func Get(shortID string) (string, bool) {
    mu.RLock()
    defer mu.RUnlock()
    url, ok := store[shortID]
    return url, ok
}