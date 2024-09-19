package testing

import (
	"Go-cache-library/cache"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRedisCache(t *testing.T) {

	cache := cache.NewRedisCache()

	key := "test:key"
	value := map[string]interface{}{"name": "Alice", "age": 30}
	ttl := 10 * time.Second
    
	t.Run("Set", func(t *testing.T) {
		err := cache.Set(key, value, ttl)
		assert.NoError(t, err, "Error setting cache")

		val, err := cache.Get(key)
		assert.NoError(t, err, "Error getting cache")

		valMap, ok := val.(map[string]interface{})
		assert.True(t, ok, "Retrieved value is not a map")

		valMap["age"] = int(valMap["age"].(float64))

		assert.Equal(t, value, valMap, "Cached value does not match")
	})

	t.Run("Get", func(t *testing.T) {
		val, err := cache.Get(key)
		assert.NoError(t, err, "Error getting cache")

		valMap, ok := val.(map[string]interface{})
		assert.True(t, ok, "Retrieved value is not a map")

		valMap["age"] = int(valMap["age"].(float64))

		assert.Equal(t, value, valMap, "Cached value does not match")
	})

	t.Run("Delete", func(t *testing.T) {
		err := cache.Delete(key)
		assert.NoError(t, err, "Error deleting cache")

		_, err = cache.Get(key)
		assert.Error(t, err, "Expected error getting deleted cache")
	})

	t.Run("Clear", func(t *testing.T) {
		cache.Set("test:key1", value, ttl)
		cache.Set("test:key2", value, ttl)

		err := cache.Clear()
		assert.NoError(t, err, "Error clearing cache")

		_, err = cache.Get("test:key1")
		assert.Error(t, err, "Expected error getting cleared cache")

		_, err = cache.Get("test:key2")
		assert.Error(t, err, "Expected error getting cleared cache")
	})

	t.Run("GetAllKeys", func(t *testing.T) {
        key1 := "test_key1"
        value1 := map[string]interface{}{"name": "Alice"}

        key2 := "test_key2"
        value2 := map[string]interface{}{"name": "Bob"}

        err := cache.Set(key1, value1, 10*time.Second)
        assert.NoError(t, err, "Error setting cache")

        err = cache.Set(key2, value2, 10*time.Second)
        assert.NoError(t, err, "Error setting cache")

        items, err := cache.GetAllKeys()
        assert.NoError(t, err, "Error getting all keys")

        assert.Contains(t, items, key1, "Key1 should be in the cache")
        assert.Contains(t, items, key2, "Key2 should be in the cache")
        assert.Equal(t, value1, items[key1], "Cached value for key1 does not match")
        assert.Equal(t, value2, items[key2], "Cached value for key2 does not match")
    })

}

