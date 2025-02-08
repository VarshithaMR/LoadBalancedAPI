package service

import "sync"

var mu sync.Mutex

func TrackUniqueRequest(id string) int {
	mu.Lock()
	defer mu.Unlock()

	var uniqueRequests = make(map[string]struct{})
	uniqueRequests[id] = struct{}{}
	return len(uniqueRequests)
}
