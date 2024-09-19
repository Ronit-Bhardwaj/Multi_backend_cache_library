package testing

import (
	"Go-cache-library/cache"
	"testing"
	"time"
    "strconv"
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

func BenchmarkRedisCacheDelete(b *testing.B) {
    cache := cache.NewRedisCache()
    _ = cache.Set("benchmark:key", "value", 10*time.Second)
    for i := 0; i < b.N; i++ {
        _ = cache.Delete("benchmark:key")
    }
}

func BenchmarkRedisCacheClear(b *testing.B) {
    cache := cache.NewRedisCache()
    _ = cache.Set("benchmark:key1", "value1", 10*time.Second)
    _ = cache.Set("benchmark:key2", "value2", 10*time.Second)
    for i := 0; i < b.N; i++ {
        _ = cache.Clear()
    }
}

func BenchmarkRedisCacheGetAllKeys(b *testing.B) {
    cache := cache.NewRedisCache()

    
    for i := 0; i < 100; i++ {
        _ = cache.Set("benchmark:key"+strconv.Itoa(i), "value", 10*time.Second)
    }

    b.ResetTimer() 
    for i := 0; i < b.N; i++ {
        _, _ = cache.GetAllKeys()
    }
}