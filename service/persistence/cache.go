package persistence

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

const redisKey = "unique_requests"

func InitRedis(host string, port int) {
	address := fmt.Sprintf("%s:%d", host, port)
	rdb = redis.NewClient(&redis.Options{
		Addr: address,
	})
}

func AddUniqueRequest(id string) error {
	_, err := rdb.SAdd(context.Background(), redisKey, id).Result()
	return err
}

func GetRedisUniqueCount() int {
	redisCount, err := rdb.SCard(context.Background(), "unique_requests").Result()
	if err != nil {
		log.Printf("Error getting Redis unique request count: %v", err)
		redisCount = 0
	}
	return int(redisCount)
}
