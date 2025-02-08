package service

import (
	"log"
	"sync"

	"LoadBalancedAPI/service/persistence"
)

var (
	mu             sync.Mutex
	uniqueRequests = make(map[string]struct{})
)

func TrackUniqueRequest(id string) int {
	mu.Lock()
	defer mu.Unlock()

	uniqueRequests[id] = struct{}{}
	err := persistence.AddUniqueRequest(id)
	if err != nil {
		log.Printf("Error adding ID to Redis: %v", err)
	}

	return len(uniqueRequests)
}

func GetUniqueRequestCount(uniqueRequestsCount int) int {
	// In-memory count
	inMemoryCount := uniqueRequestsCount

	// Redis count
	redisCount := persistence.GetRedisUniqueCount()

	// max of both counts
	return max(inMemoryCount, redisCount)
}
