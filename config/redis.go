package config

import (
	"fmt"

	"github.com/dhuki/rest-template-v2/common"
	"github.com/go-redis/redis/v8"
)

func NewRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", common.RedisHost, common.RedisPort),
		Password: "", // no password
		DB:       0,  // set default db (you can add db and become db 1, 2, 3 or etc)
	})
}
