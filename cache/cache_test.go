package cache

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// TODO: Write tests for each of the following scenarios.

func TestSetAndGet(t *testing.T) {
	// 1. Create a cache (cleanup interval doesn't matter here, just pick something)
	c := NewCache(1 * time.Second)

	// 2. Set a key with no expiration (ttl=0 means never expires)
	c.Set("name", []byte("jin"), 0)

	// 3. Get it back
	val, ok := c.Get("name")

	// 4. Verify it was found
	if !ok {
		t.Fatal("expected key 'name' to exist, but got ok=false")
	}

	// 5. Verify the value matches
	if string(val) != "jin" {
		t.Fatalf("expected value 'jin', got '%s'", string(val))
	}
}

func TestGetMiss(t *testing.T) {
	c := NewCache(1 * time.Second)

	// Try to get a key that was never set
	_, ok := c.Get("doesnotexist")

	// It should return false
	if ok {
		t.Fatal("expected ok=false for missing key, got true")
	}
}

func TestDelete(t *testing.T) {
	c := NewCache(1 * time.Second)

	// Set a key, then delete it
	c.Set("temp", []byte("data"), 0)
	c.Delete("temp")

	// Verify it's gone
	_, ok := c.Get("temp")
	if ok {
		t.Fatal("expected key 'temp' to be deleted, but it still exists")
	}
}

func TestTTLExpiration(t *testing.T) {
	// TODO: Set a key with a short TTL (e.g. 50ms), sleep past the TTL,
	// verify Get returns false.
	c := NewCache(5 * time.Second)
	c.Set("name", []byte("jin"), 5*time.Millisecond)

	time.Sleep(1 * time.Second)

	_, ok := c.Get("name")

	if ok {
		t.Fatal("expected ok=false for passing ttl, got ture")
	}
}

func TestEviction(t *testing.T) {
	// TODO: Create a cache with a short cleanup interval,
	// set items with short TTL, wait for eviction, verify Count() drops.
	c := NewCache(5 * time.Millisecond)
	c.Set("name", []byte("jin"), 1*time.Millisecond)
	old_cnt := c.Count()

	time.Sleep(1 * time.Second)
	new_cnt := c.Count()

	if new_cnt >= old_cnt {
		t.Fatal("expected the count to fall down, but it goes up")
	}

}

func TestConcurrentAccess(t *testing.T) {
	c := NewCache(1 * time.Second)

	// WaitGroup tracks how many goroutines are still running.
	// Think of it as a counter:
	//   wg.Add(1)  → counter++  (call BEFORE launching a goroutine)
	//   wg.Done()  → counter--  (call when a goroutine finishes)
	//   wg.Wait()  → blocks until counter reaches 0
	var wg sync.WaitGroup

	// Launch 10 goroutines, each doing 100 operations
	for i := 0; i < 10; i++ {
		wg.Add(1) // tell WaitGroup: "one more goroutine to wait for"

		go func(id int) {
			defer wg.Done() // when this goroutine exits, decrement the counter

			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("key-%d-%d", id, j)
				c.Set(key, []byte("val"), 5*time.Second)
				c.Get(key)
				c.Delete(key)
			}
		}(i)
	}

	// Block here until all 10 goroutines have called wg.Done()
	wg.Wait()

	// If we reach here without a panic or race detector complaint, the test passes.
}
