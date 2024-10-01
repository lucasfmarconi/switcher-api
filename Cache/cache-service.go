package CacheServices

import (
	"context"
	"fmt"
	"time"

	"github.com/allegro/bigcache/v3"
)

func SetCache(key, value string) {
	cache, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(10 * time.Minute))

	cacheSetResult := cache.Set(key, []byte(value))
	if cacheSetResult != nil {
		fmt.Println(string("Error setting cache: " + cacheSetResult.Error()))
		return 
	}

	entry, _ := cache.Get(key)
	fmt.Println(string(entry))
}

func GetCache(key string) string {
	cache, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(10 * time.Minute))

	entry, _ := cache.Get(key)
	return string(entry)
}