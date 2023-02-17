package config

import (
	"casbin-gin/cmd/runtime"
	"casbin-gin/common/component/storage/cache"
	"log"
)

type Cache struct {
	Redis  *RedisConnectOptions
	Memory interface{}
}

var CacheConfig = new(Cache)

func (e Cache) Setup() {
	if e.Redis != nil {
		options, err := e.Redis.GetRedisOptions()
		if err != nil {
			log.Fatal("Config cache setup failed, err:", err)
		}
		r, err := cache.NewRedis(GetRedisClient(), options)
		if err != nil {
			log.Fatal("Config cache setup failed, err:", err)
		}
		if _redis == nil {
			_redis = r.GetClient()
		}
		//return r, nil
		runtime.ApplicationContext.SetCacheAdapter(r)
		return
	}
	m := cache.NewMemory()
	runtime.ApplicationContext.SetCacheAdapter(m)
}
