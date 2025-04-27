package v4

// ActuatorLinks represents the available actuator endpoints
type ActuatorLinks struct {
	Links map[string]ActuatorLink `json:"_links"`
}

// ActuatorLink represents a single link in the actuator API
type ActuatorLink struct {
	Href      string `json:"href"`
	Templated bool   `json:"templated"`
}

// ActuatorVersionRoute represents a version and its associated routes
type ActuatorVersionRoute struct {
	Version string   `json:"version"`
	Routes  []string `json:"routes"`
}

// ActuatorNamespaceVersionRoutes represents a namespace and its version routes
type ActuatorNamespaceVersionRoutes struct {
	Namespace     string                 `json:"namespace"`
	VersionRoutes []ActuatorVersionRoute `json:"versionRoutes"`
}

// ActuatorRateLimit represents rate limit information for an API endpoint
type ActuatorRateLimit struct {
	Count     int    `json:"count"`
	TimeUnits string `json:"timeUnits"`
}

// ActuatorRateLimitConfig represents a rate limit configuration for an API endpoint
type ActuatorRateLimitConfig struct {
	Method    string            `json:"method"`
	Path      string            `json:"path"`
	RateLimit ActuatorRateLimit `json:"rateLimit"`
}

// ActuatorCircuitBreaker represents the state of a circuit breaker
type ActuatorCircuitBreaker struct {
	CircuitBreaker string `json:"circuitbreaker"`
	State          string `json:"state"`
}

// ActuatorCachesResponse represents the response from the caches endpoint
type ActuatorCachesResponse struct {
	CacheManagers map[string]ActuatorCacheManager `json:"cacheManagers"`
}

// ActuatorCacheManager represents a cache manager in the caches response
type ActuatorCacheManager struct {
	Caches map[string]ActuatorCache `json:"caches"`
}

// ActuatorCache represents a single cache in the caches response
type ActuatorCache struct {
	Target string `json:"target"`
}

// ActuatorCacheResponse represents the response for a specific cache
type ActuatorCacheResponse struct {
	Name   string `json:"name"`
	Target string `json:"target"`
}