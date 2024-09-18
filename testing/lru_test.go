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

    lru.Set("1", 100, 10*time.Second)
    lru.Set("2", 200, 10*time.Second)
    lru.Set("3", 300, 10*time.Second)

    keys := lru.GetAllKeys()
    if len(keys) != 3 {
        t.Errorf("expected 3 keys, got %d", len(keys))
    }

    if !contains(keys, "1") || !contains(keys, "2") || !contains(keys, "3") {
        t.Errorf("keys do not contain expected values")
    }
}

func contains(slice []string, item string) bool {
    for _, a := range slice {
        if a == item {
            return true
        }
    }
    return false
}