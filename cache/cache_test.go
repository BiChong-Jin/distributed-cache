package cache

import (
	"testing"
	"time"
)

// TODO: Write tests for each of the following scenarios.

func TestSetAndGet(t *testing.T) {
	// TODO: Create a cache, set a key, get it back, verify the value matches.
}

func TestGetMiss(t *testing.T) {
	// TODO: Get a key that was never set, verify it returns false.
}

func TestDelete(t *testing.T) {
	// TODO: Set a key, delete it, verify Get returns false.
}

func TestTTLExpiration(t *testing.T) {
	// TODO: Set a key with a short TTL (e.g. 50ms), sleep past the TTL,
	// verify Get returns false.
}

func TestEviction(t *testing.T) {
	// TODO: Create a cache with a short cleanup interval,
	// set items with short TTL, wait for eviction, verify Count() drops.
}

func TestConcurrentAccess(t *testing.T) {
	// TODO: Launch multiple goroutines doing Set/Get/Delete simultaneously.
	// Use sync.WaitGroup. The test passes if there's no race condition (run with -race).
}
