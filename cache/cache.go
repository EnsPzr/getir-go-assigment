package cache

import (
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

var Cache CacheService

type CacheService struct {
	client *cache.Cache
}

// InitGoCache
// This function initialize Cache variable.
// Run after main.go
func InitGoCache() {
	log.Println("Cache init started")
	Cache.client = cache.New(5*time.Minute, 10*time.Minute)
	log.Println("Cache init finished")
}

// Get
// This method gets value inside cache according to key.
// Method return two variable. 1- Value, 2- If key is exist
func (c *CacheService) Get(key string) (interface{}, bool) {
	val, ok := c.client.Get(key)
	return val, ok
}

// Set
// This method holds value in cache with key.
// Method no return.
func (c *CacheService) Set(key string, value interface{}) {
	c.client.Set(key, value, 5*time.Minute)
}
