package main

import (
	"Go-cache-library/cache"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var redisCache *cache.RedisCache
var lruCache *cache.LRU

func init() {
	lruCache = cache.Newlru(100)
	redisCache = cache.NewRedisCache()
}

type RequestBody struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
	TTL   int         `json:"ttl"`
}

func main() {
	router := gin.Default()

	router.POST("/set", func(c *gin.Context) {
		var body RequestBody
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ttl := time.Duration(body.TTL) * time.Second

		if err := lruCache.Set(body.Key, body.Value, ttl); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := redisCache.Set(body.Key, body.Value, ttl); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "set successful"})
	})

	router.GET("/get/:key", func(c *gin.Context) {
		key := c.Param("key")

		if val, err := lruCache.Get(key); err == nil {
			c.JSON(http.StatusOK, gin.H{"value": val, "source": "LRU"})
			return
		}

		if val, err := redisCache.Get(key); err == nil {
			c.JSON(http.StatusOK, gin.H{"value": val, "source": "Redis"})
			return
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "key not found"})
	})

	router.GET("/get", func(c *gin.Context) {
		
		lruItems := lruCache.GetAllKeys()
		
		allKeys, err := redisCache.GetAllKeys() 
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{
			"lru_items":   lruItems,
			"keys": allKeys,
		})
	})

	router.DELETE("/delete/:key", func(c *gin.Context) {
		key := c.Param("key")

		if err := lruCache.Delete(key); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if err := redisCache.Delete(key); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "delete successful"})
	})

	router.POST("/clear", func(c *gin.Context) {
		lruCache.Clear()
		if err := redisCache.Clear(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "cache cleared"})
	})

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	go func() {
		<-stopChan
		lruCache.StopCleanup()
		os.Exit(0)
	}()

	log.Fatal(router.Run(":8080"))
}
