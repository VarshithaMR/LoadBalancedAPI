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

	// test connection
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	} else {
		log.Printf("Successfully connected to Redis on: %v", port)
	}
}

func AddUniqueRequest(id string) error {
	_, err := rdb.SAdd(context.Background(), redisKey, id).Result()
	if err == nil {
		log.Printf("Added unique id to redis: %v", id)
	}
	members, err := rdb.SMembers(context.Background(), redisKey).Result()
	if err != nil {
		log.Printf("Error fetching Redis set members: %v", err)
	}
	log.Printf("Current Redis set members: %v", members)
	return err
}

func GetRedisUniqueCount() int {
	redisCount, err := rdb.SCard(context.Background(), redisKey).Result()
	if err != nil {
		log.Printf("Error getting Redis unique request count: %v", err)
		redisCount = 0
	}
	log.Printf("Redis Unique count: %v", int(redisCount))
	return int(redisCount)
}
