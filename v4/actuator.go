package v4

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/nutanix-cloud-native/prism-go-client"
	"github.com/nutanix-cloud-native/prism-go-client/internal"
)

// ActuatorAPI manages the actuator API
type ActuatorAPI struct {
	client *internal.Client
}

// NewActuatorAPI returns a new ActuatorAPI instance
func NewActuatorAPI(credentials prismgoclient.Credentials, opts ...internal.ClientOption) (*ActuatorAPI, error) {
	clientOpts := []internal.ClientOption{
		internal.WithCredentials(&credentials),
		internal.WithUserAgent("prism-go-client"),
		internal.WithAbsolutePath("/api/actuator"),
	}

	// Add any additional options
	clientOpts = append(clientOpts, opts...)

	client, err := internal.NewClient(clientOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create actuator client: %v", err)
	}

	return &ActuatorAPI{
		client: client,
	}, nil
}

// GetVersionRoutes retrieves all version routes from the actuator API
func (a *ActuatorAPI) GetVersionRoutes(ctx context.Context) ([]ActuatorNamespaceVersionRoutes, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/versionroutes", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var versionRoutes []ActuatorNamespaceVersionRoutes
	err = a.client.Do(ctx, req, &versionRoutes)
	if err != nil {
		return nil, fmt.Errorf("error getting version routes: %v", err)
	}

	return versionRoutes, nil
}

// GetRateLimits retrieves rate limits information from the actuator API
func (a *ActuatorAPI) GetRateLimits(ctx context.Context) ([]ActuatorRateLimitConfig, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/rate-limits", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var rateLimits []ActuatorRateLimitConfig
	err = a.client.Do(ctx, req, &rateLimits)
	if err != nil {
		return nil, fmt.Errorf("error getting rate limits: %v", err)
	}

	return rateLimits, nil
}

// GetCircuitBreakers retrieves circuit breakers information from the actuator API
func (a *ActuatorAPI) GetCircuitBreakers(ctx context.Context) ([]ActuatorCircuitBreaker, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/circuit-breakers", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var circuitBreakers []ActuatorCircuitBreaker
	err = a.client.Do(ctx, req, &circuitBreakers)
	if err != nil {
		return nil, fmt.Errorf("error getting circuit breakers: %v", err)
	}

	return circuitBreakers, nil
}

// GetBeans retrieves beans information from the actuator API
func (a *ActuatorAPI) GetBeans(ctx context.Context) (interface{}, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/beans", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var beans interface{}
	err = a.client.Do(ctx, req, &beans)
	if err != nil {
		return nil, fmt.Errorf("error getting beans: %v", err)
	}

	return beans, nil
}

// GetCaches retrieves all caches information from the actuator API
func (a *ActuatorAPI) GetCaches(ctx context.Context) (*ActuatorCachesResponse, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/caches", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var caches ActuatorCachesResponse
	err = a.client.Do(ctx, req, &caches)
	if err != nil {
		return nil, fmt.Errorf("error getting caches: %v", err)
	}

	return &caches, nil
}

// GetCacheByName retrieves information for a specific cache from the actuator API
func (a *ActuatorAPI) GetCacheByName(ctx context.Context, cacheName string) (*ActuatorCacheResponse, error) {
	endpoint := path.Join("/caches", cacheName)
	req, err := a.client.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var cache ActuatorCacheResponse
	err = a.client.Do(ctx, req, &cache)
	if err != nil {
		return nil, fmt.Errorf("error getting cache %s: %v", cacheName, err)
	}

	return &cache, nil
}

// GetHealth retrieves health information from the actuator API
func (a *ActuatorAPI) GetHealth(ctx context.Context) (interface{}, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var health interface{}
	err = a.client.Do(ctx, req, &health)
	if err != nil {
		return nil, fmt.Errorf("error getting health: %v", err)
	}

	return health, nil
}

// GetHealthByPath retrieves health information for a specific path from the actuator API
func (a *ActuatorAPI) GetHealthByPath(ctx context.Context, healthPath string) (interface{}, error) {
	endpoint := path.Join("/health", healthPath)
	req, err := a.client.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var health interface{}
	err = a.client.Do(ctx, req, &health)
	if err != nil {
		return nil, fmt.Errorf("error getting health for path %s: %v", healthPath, err)
	}

	return health, nil
}

// GetInfo retrieves info from the actuator API
func (a *ActuatorAPI) GetInfo(ctx context.Context) (interface{}, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/info", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var info interface{}
	err = a.client.Do(ctx, req, &info)
	if err != nil {
		return nil, fmt.Errorf("error getting info: %v", err)
	}

	return info, nil
}

// GetConditions retrieves conditions information from the actuator API
func (a *ActuatorAPI) GetConditions(ctx context.Context) (interface{}, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/conditions", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var conditions interface{}
	err = a.client.Do(ctx, req, &conditions)
	if err != nil {
		return nil, fmt.Errorf("error getting conditions: %v", err)
	}

	return conditions, nil
}

// GetConfigProps retrieves all configuration properties from the actuator API
func (a *ActuatorAPI) GetConfigProps(ctx context.Context) (interface{}, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/configprops", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var configProps interface{}
	err = a.client.Do(ctx, req, &configProps)
	if err != nil {
		return nil, fmt.Errorf("error getting config properties: %v", err)
	}

	return configProps, nil
}

// GetConfigPropsByPrefix retrieves configuration properties by prefix from the actuator API
func (a *ActuatorAPI) GetConfigPropsByPrefix(ctx context.Context, prefix string) (interface{}, error) {
	endpoint := path.Join("/configprops", prefix)
	req, err := a.client.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var configProps interface{}
	err = a.client.Do(ctx, req, &configProps)
	if err != nil {
		return nil, fmt.Errorf("error getting config properties for prefix %s: %v", prefix, err)
	}

	return configProps, nil
}

// GetEnv retrieves environment information from the actuator API
func (a *ActuatorAPI) GetEnv(ctx context.Context) (interface{}, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/env", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var env interface{}
	err = a.client.Do(ctx, req, &env)
	if err != nil {
		return nil, fmt.Errorf("error getting environment info: %v", err)
	}

	return env, nil
}

// GetEnvByMatch retrieves environment information by match pattern from the actuator API
func (a *ActuatorAPI) GetEnvByMatch(ctx context.Context, toMatch string) (interface{}, error) {
	endpoint := path.Join("/env", toMatch)
	req, err := a.client.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var env interface{}
	err = a.client.Do(ctx, req, &env)
	if err != nil {
		return nil, fmt.Errorf("error getting environment info for pattern %s: %v", toMatch, err)
	}

	return env, nil
}

// GetLoggers retrieves all loggers information from the actuator API
func (a *ActuatorAPI) GetLoggers(ctx context.Context) (interface{}, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/loggers", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var loggers interface{}
	err = a.client.Do(ctx, req, &loggers)
	if err != nil {
		return nil, fmt.Errorf("error getting loggers: %v", err)
	}

	return loggers, nil
}

// GetLoggerByName retrieves logger information by name from the actuator API
func (a *ActuatorAPI) GetLoggerByName(ctx context.Context, name string) (interface{}, error) {
	endpoint := path.Join("/loggers", name)
	req, err := a.client.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var logger interface{}
	err = a.client.Do(ctx, req, &logger)
	if err != nil {
		return nil, fmt.Errorf("error getting logger %s: %v", name, err)
	}

	return logger, nil
}

// GetHeapDump retrieves heap dump from the actuator API
func (a *ActuatorAPI) GetHeapDump(ctx context.Context) (interface{}, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/heapdump", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var heapDump interface{}
	err = a.client.Do(ctx, req, &heapDump)
	if err != nil {
		return nil, fmt.Errorf("error getting heap dump: %v", err)
	}

	return heapDump, nil
}

// GetThreadDump retrieves thread dump from the actuator API
func (a *ActuatorAPI) GetThreadDump(ctx context.Context) (interface{}, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/threaddump", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var threadDump interface{}
	err = a.client.Do(ctx, req, &threadDump)
	if err != nil {
		return nil, fmt.Errorf("error getting thread dump: %v", err)
	}

	return threadDump, nil
}

// GetMetrics retrieves all metrics information from the actuator API
func (a *ActuatorAPI) GetMetrics(ctx context.Context) (interface{}, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/metrics", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var metrics interface{}
	err = a.client.Do(ctx, req, &metrics)
	if err != nil {
		return nil, fmt.Errorf("error getting metrics: %v", err)
	}

	return metrics, nil
}

// GetMetricByName retrieves metric information by name from the actuator API
func (a *ActuatorAPI) GetMetricByName(ctx context.Context, metricName string) (interface{}, error) {
	endpoint := path.Join("/metrics", metricName)
	req, err := a.client.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var metric interface{}
	err = a.client.Do(ctx, req, &metric)
	if err != nil {
		return nil, fmt.Errorf("error getting metric %s: %v", metricName, err)
	}

	return metric, nil
}

// GetScheduledTasks retrieves scheduled tasks information from the actuator API
func (a *ActuatorAPI) GetScheduledTasks(ctx context.Context) (interface{}, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/scheduledtasks", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var scheduledTasks interface{}
	err = a.client.Do(ctx, req, &scheduledTasks)
	if err != nil {
		return nil, fmt.Errorf("error getting scheduled tasks: %v", err)
	}

	return scheduledTasks, nil
}

// GetStartup retrieves startup information from the actuator API
func (a *ActuatorAPI) GetStartup(ctx context.Context) (interface{}, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/startup", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var startup interface{}
	err = a.client.Do(ctx, req, &startup)
	if err != nil {
		return nil, fmt.Errorf("error getting startup info: %v", err)
	}

	return startup, nil
}

// GetMappings retrieves mappings information from the actuator API
func (a *ActuatorAPI) GetMappings(ctx context.Context) (interface{}, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/mappings", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var mappings interface{}
	err = a.client.Do(ctx, req, &mappings)
	if err != nil {
		return nil, fmt.Errorf("error getting mappings: %v", err)
	}

	return mappings, nil
}

// RefreshActuator triggers a refresh of the actuator API
func (a *ActuatorAPI) RefreshActuator(ctx context.Context) (interface{}, error) {
	req, err := a.client.NewRequest(http.MethodPost, "/refresh", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var result interface{}
	err = a.client.Do(ctx, req, &result)
	if err != nil {
		return nil, fmt.Errorf("error refreshing actuator: %v", err)
	}

	return result, nil
}

// GetFeatures retrieves features information from the actuator API
func (a *ActuatorAPI) GetFeatures(ctx context.Context) (interface{}, error) {
	req, err := a.client.NewRequest(http.MethodGet, "/features", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var features interface{}
	err = a.client.Do(ctx, req, &features)
	if err != nil {
		return nil, fmt.Errorf("error getting features: %v", err)
	}

	return features, nil
}

// GetActuatorEndpoints retrieves all available actuator endpoints
func (a *ActuatorAPI) GetActuatorEndpoints(ctx context.Context) (*ActuatorLinks, error) {
	req, err := a.client.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	var links ActuatorLinks
	err = a.client.Do(ctx, req, &links)
	if err != nil {
		return nil, fmt.Errorf("error getting actuator endpoints: %v", err)
	}

	return &links, nil
}