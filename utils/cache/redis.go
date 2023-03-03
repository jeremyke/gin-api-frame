/**
* @Author: maxwell
* @description:redis封装
* @Date: 2021/9/22 8:36 下午
 */
package cache

import (
	"context"
	"gin-api-frame/app/global/variable"
	"gin-api-frame/utils"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	BillingRedis *Redis
	DataRedis    *Redis
)

type Redis struct {
	client *redis.Client
}

func InitRedis() {
	BillingRedis = newRedis("REDIS_RW")
	DataRedis = newRedis("REDIS_RW")
}

func newRedis(connectName string) *Redis {
	conf := getConnectConf(connectName)
	c := redis.NewClient(conf)
	if err := c.Ping(context.Background()).Err(); err != nil {

		panic(connectName + " failed to connect to Redis server")
	}
	return &Redis{client: c}
}

func getConnectConf(connectName string) *redis.Options {
	return &redis.Options{
		Addr:     utils.HostPort(variable.EnvConfig.GetString(connectName+"_HOST"), variable.EnvConfig.GetString(connectName+"_PORT")),
		Password: variable.EnvConfig.GetString(connectName + "_PASSWORD"),
		DB:       13,
	}
}

func (r *Redis) Cli() *redis.Client {
	return r.client
}

func (r *Redis) GetKey(key string) (string, error) {
	return r.client.Get(context.Background(), key).Result()
}

func (r *Redis) SetKey(key, value string, expire time.Duration) error {
	return r.client.Set(context.Background(), key, value, expire*time.Second).Err()
}

func (r *Redis) KeyExist(key string) (exist bool, err error) {
	count, err := r.client.Exists(context.Background(), key).Result()
	if err != nil {
		return
	}

	if count < 1 {
		return
	}

	exist = true
	return
}

func (r *Redis) HExists(key, field string) (exist bool, err error) {
	exist, err = r.client.HExists(context.Background(), key, field).Result()
	if err != nil {
		return
	}
	return
}

func (r *Redis) HGet(key, field string) (string, error) {
	return r.client.HGet(context.Background(), key, field).Result()
}

func (r *Redis) HSet(key string, value ...interface{}) error {
	return r.client.HSet(context.Background(), key, value).Err()
}

func (r *Redis) HGetAll(key string) (m map[string]string, err error) {
	return r.client.HGetAll(context.Background(), key).Result()
}

// Select
func (r *Redis) Select(database int) error {
	return r.client.Pipeline().Select(context.Background(), database).Err()
}

func (r *Redis) GetPipe() redis.Pipeliner {
	return r.client.Pipeline()
}

// CheckTTL
func (r *Redis) CheckTTL(key string) (ret time.Duration, err error) {
	emptyCtx := context.Background()
	return r.client.TTL(emptyCtx, key).Result()
}

// SetNX
func (r *Redis) SetNX(key, value string, expire time.Duration) (bool, error) {
	emptyCtx := context.Background()
	return r.client.SetNX(emptyCtx, key, value, expire).Result()
}
