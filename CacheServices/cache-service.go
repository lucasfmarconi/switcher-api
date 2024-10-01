package CacheServices

import (
	"context"
	"fmt"
	"time"

	"github.com/allegro/bigcache/v3"
)

//Define a global property to hold the cache instance
var globalcache *bigcache.BigCache

func init() {
    // Initialize the cache and assign it to the global variable
	var err error
	cache, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
        panic(err)
    }
	globalcache = cache
}

func SetCache(key, value string) {

	cacheSetResult := globalcache.Set(key, []byte(value))
	if cacheSetResult != nil {
		fmt.Println(string("Error setting cache: " + cacheSetResult.Error()))
		return 
	}
	fmt.Println(key + " - cached successfully")
}

func GetCache(key string) string {
	entry, _ := globalcache.Get(key)
	return string(entry)
}