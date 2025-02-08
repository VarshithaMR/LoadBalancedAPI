# Tech-Stack
* go version - 1.23.3
* IDE : Intellij IDEA
* Postman
* RedisCache Library - "github.com/go-redis/redis/v8"

# FAQ
* API - GET http://localhost:8080/api/verve/accept
* Redis installation & testing:
  brew install redis
  brew services start redis
  redis-cli ping

# Design implementation
* Application exposes a single GET API with id as mandatory query parameter, endpoint as optional query parameter. When the optional query parameter is present, a POST request for same endpoint query param which is not exposed is triggered.
    Responses:
        - "200 Ok"
        - "400 BadRequest" - ID parameter is required
* Request handling based on id and endpoint and also tracks the unique id request counts using cache and in-memory store.
* Concurrency handling using goroutines to handle 10K incoming requests.
* Each unique id will be tracked using an in-memory data structure such as a map or set.
* Singleton pattern for cache connection.
* Mutexes used to ensure that updates to shared resources  are done in a thread-safe manner.
* Redis operations (like SADD, SMembers and SCARD) will be used to add ids to a Redis set and count the unique ids efficiently.
* Sophisticated error handling.
* Logs are generated for every incoming request to track errors and status. 
* Periodic logs of the unique requests count.
* Externalised configuration