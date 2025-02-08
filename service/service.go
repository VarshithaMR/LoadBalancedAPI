package service

func TrackUniqueRequest(id string) int {
	var uniqueRequests = make(map[string]struct{})
	uniqueRequests[id] = struct{}{}
	return len(uniqueRequests)
}
