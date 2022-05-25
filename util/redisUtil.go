package util

import (
  "context"
  "github.com/go-redis/redis/v8"
)

var (
  rdb *redis.Client
  ctx context.Context
)

func InitRedis() {
  ctx = context.Background()
  rdb = redis.NewClient(&redis.Options{
          Addr:	  "localhost:13579",
          Password: "", // no password set
          DB:		  0,  // use default DB
  })
}

func HGetAll(key string) map[string]string {
  machinesStatus, _ := rdb.HGetAll(ctx, key).Result()
  return machinesStatus
}

func HSet(key string, feild string, value string) {
  rdb.HSet(ctx, key, feild, value)
}
