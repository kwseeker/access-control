package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var defaultCtx = context.Background()

// Redis cache implement
type Redis struct {
	client *redis.Client
}

// NewRedis redis模式
func NewRedis(client *redis.Client, options *redis.Options) (*Redis, error) {
	if client == nil {
		client = redis.NewClient(options)
	}
	r := &Redis{
		client: client,
	}
	err := r.connect()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (*Redis) String() string {
	return "redis"
}

// connect connect test
func (r *Redis) connect() error {
	var err error
	_, err = r.client.Ping(defaultCtx).Result()
	return err
}

// Get from key
func (r *Redis) Get(key string) (string, error) {
	//return r.client.Get(key).Result()
	return r.client.Get(defaultCtx, key).Result()
}

// Set value with key and expire time
func (r *Redis) Set(key string, val interface{}, expire int) error {
	return r.client.Set(defaultCtx, key, val, time.Duration(expire)*time.Second).Err()
}

// Del delete key in redis
func (r *Redis) Del(key string) error {
	return r.client.Del(defaultCtx, key).Err()
}

// HashGet from key
func (r *Redis) HashGet(hk, key string) (string, error) {
	return r.client.HGet(defaultCtx, hk, key).Result()
}

// HashDel delete key in specify redis's hashtable
func (r *Redis) HashDel(hk, key string) error {
	return r.client.HDel(defaultCtx, hk, key).Err()
}

func (r *Redis) Increase(key string) error {
	return r.client.Incr(defaultCtx, key).Err()
}

func (r *Redis) Decrease(key string) error {
	return r.client.Decr(defaultCtx, key).Err()
}

func (r *Redis) Expire(key string, dur time.Duration) error {
	return r.client.Expire(defaultCtx, key, dur).Err()
}

// GetClient 暴露原生client
func (r *Redis) GetClient() *redis.Client {
	return r.client
}
