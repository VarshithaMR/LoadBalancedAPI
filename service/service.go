package service

import (
	"log"
	"sync"
	"time"

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

func LogUniqueRequestsEveryMinute() {
	// setting up the ticker
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		<-ticker.C
		mu.Lock()
		count := len(uniqueRequests)
		log.Printf("Unique requests in the last minute: %d", count)
		mu.Unlock()
	}
}
