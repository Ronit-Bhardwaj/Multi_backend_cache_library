package testing

import (
	"Go-cache-library/cache"
	"testing"
	"time"
)

func TestLRU_SetAndGet(t *testing.T) {
	lru := cache.Newlru(3)

	// Testing set and get
	lru.Set("1", 100, 10*time.Second)
	val, err := lru.Get("1")
	if err != nil || val != 100 {
		t.Errorf("expected 100, got %v", val)
	}

	// Test expiration
	lru.Set("2", 200, 10*time.Second)
	time.Sleep(11 * time.Second)
	_, err = lru.Get("2")
	if err == nil {
		t.Errorf("expected cache miss due to expiration")
	}

	// Testing eviction
	lru.Set("3", 300, 10*time.Second)
	lru.Set("4", 400, 10*time.Second)

	_, err = lru.Get("1")
	if err == nil {
		t.Errorf("expected eviction of key '1'")
	}

	val, err = lru.Get("4")
	if err != nil || val != 400 {
		t.Errorf("expected 400, got %v", val)
	}
}

func TestLRU_Delete(t *testing.T) {
	lru := cache.Newlru(2)

	lru.Set("1", 100, 10*time.Second)
	lru.Delete("1")

	_, err := lru.Get("1")
	if err == nil {
		t.Errorf("expected cache miss after deletion")
	}
}

func TestLRU_Clear(t *testing.T) {
	lru := cache.Newlru(2)

	lru.Set("1", 100, 10*time.Second)
	lru.Set("2", 200, 10*time.Second)
	lru.Clear()

	_, err := lru.Get("1")
	if err == nil {
		t.Errorf("expected cache miss after clearing cache")
	}

	_, err = lru.Get("2")
	if err == nil {
		t.Errorf("expected cache miss after clearing cache")
	}
}

func TestLRU_GetAllKeys(t *testing.T) {
    lru := cache.Newlru(3)

    // Set some items
    lru.Set("1", 100, 10*time.Second)
    lru.Set("2", 200, 10*time.Second)
    lru.Set("3", 300, 10*time.Second)

    // Retrieve all keys
    items := lru.GetAllKeys()
    if len(items) != 3 {
        t.Errorf("expected 3 items, got %d", len(items))
    }

    // Verify the values
    if val, ok := items["1"]; !ok || val != 100 {
        t.Errorf("expected key '1' to have value 100, got %v", val)
    }
    if val, ok := items["2"]; !ok || val != 200 {
        t.Errorf("expected key '2' to have value 200, got %v", val)
    }
    if val, ok := items["3"]; !ok || val != 300 {
        t.Errorf("expected key '3' to have value 300, got %v", val)
    }

    // Test eviction
    lru.Set("4", 400, 10*time.Second)
    lru.Set("5", 500, 10*time.Second)

    items = lru.GetAllKeys()
    if len(items) != 3 {
        t.Errorf("expected 3 items after eviction, got %d", len(items))
    }

    // Verify that evicted keys are not present
    if _, ok := items["1"]; ok {
        t.Errorf("expected key '1' to be evicted, but it was found")
    }
    if val, ok := items["4"]; !ok || val != 400 {
        t.Errorf("expected key '4' to have value 400, got %v", val)
    }
    if val, ok := items["5"]; !ok || val != 500 {
        t.Errorf("expected key '5' to have value 500, got %v", val)
    }
}
