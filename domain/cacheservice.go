package domain

import (
	"time"

	cache "github.com/patrickmn/go-cache"
)

var GlobalCache *cache.Cache = cache.New(10*time.Minute, 10*time.Minute)

func InsertCache(k string, x interface{}, d time.Duration) {
	GlobalCache.Set(k, x, d)
}

func FindCache(k string) (interface{}, bool) {
	if x, found := GlobalCache.Get(k); found {
		return x, true
	} else {
		return nil, false
	}
}
