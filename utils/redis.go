package utils

import (
	"github.com/redis/go-redis/v9"
)

var RedisClient = redis.NewClient(&redis.Options{

	Addr:     "redis:6379",
	Password: "",
	DB:       0,
})
