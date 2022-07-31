package factory

import (
	"fmt"
	"testing"
)

func TestRedisCache(t *testing.T) {
	var redisCacheF CacheFactory
	redisCacheF = RedisCacheFactory{}
	redisCache := redisCacheF.Create()
	redisCache.Set("k1", 1)
	v, _ := redisCache.Get("k1")
	res, _ := v.(int)
	fmt.Println(res)
}
