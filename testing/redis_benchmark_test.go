package testing

import (
    "Go-cache-library/cache"
    "testing"
    "time"
)

func BenchmarkRedisCacheSet(b *testing.B) {
    cache := cache.NewRedisCache()
    for i := 0; i < b.N; i++ {
        _ = cache.Set("benchmark:key", "value", 10*time.Second)
    }
}

func BenchmarkRedisCacheGet(b *testing.B) {
    cache := cache.NewRedisCache()
    _ = cache.Set("benchmark:key", "value", 10*time.Second)
    for i := 0; i < b.N; i++ {
        _, _ = cache.Get("benchmark:key")
    }
}

func BenchmarkRedisCacheGetAllKeys(b *testing.B) {
    cache := cache.NewRedisCache()
    for i := 0; i < b.N; i++ {
        _ = cache.Set("benchmark:key", "value", 10*time.Second)
        _, _ = cache.GetAllKeys()
    }
}