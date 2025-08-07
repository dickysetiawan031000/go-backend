package utils

import (
	"sync"
	"time"
)

var (
	blacklistedTokens = make(map[string]time.Time)
	mu                sync.Mutex
)

// BlacklistToken menambahkan token ke daftar blacklist
func BlacklistToken(token string, expiration time.Time) {
	mu.Lock()
	defer mu.Unlock()
	blacklistedTokens[token] = expiration
}

// IsTokenBlacklisted mengecek apakah token ada di blacklist
func IsTokenBlacklisted(token string) bool {
	mu.Lock()
	defer mu.Unlock()

	exp, exists := blacklistedTokens[token]
	if !exists {
		return false
	}

	// Hapus token jika sudah expired
	if time.Now().After(exp) {
		delete(blacklistedTokens, token)
		return false
	}

	return true
}
