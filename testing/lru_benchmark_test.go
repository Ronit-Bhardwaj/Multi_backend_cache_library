package testing

import (
    "Go-cache-library/cache"
    "testing"
    "time"
    "fmt"
)
func BenchmarkLRUSet(b *testing.B) {
    lru := cache.Newlru(100)
    for i := 0; i < b.N; i++ {
        _ = lru.Set("benchmark:key", "value", 10*time.Second)
    }
}

func BenchmarkLRUGet(b *testing.B) {
    lru := cache.Newlru(100)
    _ = lru.Set("benchmark:key", "value", 10*time.Second)
    for i := 0; i < b.N; i++ {
        _, _ = lru.Get("benchmark:key")
    }
}
func BenchmarkLRUDelete(b *testing.B) {
    lru := cache.Newlru(100)
    _ = lru.Set("benchmark:key", "value", 10*time.Second)
    for i := 0; i < b.N; i++ {
        _ = lru.Delete("benchmark:key")
    }
}

func BenchmarkLRUClear(b *testing.B) {
    lru := cache.Newlru(100)
    _ = lru.Set("benchmark:key1", "value1", 10*time.Second)
    _ = lru.Set("benchmark:key2", "value2", 10*time.Second)
    for i := 0; i < b.N; i++ {
        lru.Clear()
    }
}

func BenchmarkLRUGetAllKeys(b *testing.B) {
    lru := cache.Newlru(100)
    for i := 0; i < 100; i++ {
        _ = lru.Set(fmt.Sprintf("key%d", i), i, 10*time.Second)
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = lru.GetAllKeys()
    }
}
