package fflags

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"
)

// FFlags represents the main feature flags service
type FFlags struct {
	rand *rand.Rand
}

// NewFFlags creates a new FFlags instance
func NewFFlags() *FFlags {
	return &FFlags{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// ConfigService interface for configuration service
type ConfigService interface {
	BoolValue(ctx context.Context, key string, params interface{}) (bool, error)
	StringValue(ctx context.Context, key string, params interface{}) (string, error)
	IntValue(ctx context.Context, key string, params interface{}) (int, error)
	Int32Value(ctx context.Context, key string, params interface{}) (int32, error)
	Float64Value(ctx context.Context, key string, params interface{}) (float64, error)
	JsonValue(ctx context.Context, key string, params interface{}) (json.RawMessage, error)
}

// MockConfigService implements ConfigService
type MockConfigService struct {
	client interface{}
	logger interface{}
	config interface{}
	rand   *rand.Rand
}

// NewConfigService creates a new config service
func NewConfigService(client interface{}, logger interface{}, config interface{}) ConfigService {
	return &MockConfigService{
		client: client,
		logger: logger,
		config: config,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// NewEventsConfigService creates a new events config service
func NewEventsConfigService(client interface{}, logger interface{}, config interface{}) ConfigService {
	return &MockConfigService{
		client: client,
		logger: logger,
		config: config,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// NewConnectConfigService creates a new connect config service
func NewConnectConfigService(client interface{}, logger interface{}, config interface{}) ConfigService {
	return &MockConfigService{
		client: client,
		logger: logger,
		config: config,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// BoolValue returns a random boolean
func (m *MockConfigService) BoolValue(ctx context.Context, key string, params interface{}) (bool, error) {
	return m.rand.Intn(2) == 1, nil
}

// StringValue returns a random string
func (m *MockConfigService) StringValue(ctx context.Context, key string, params interface{}) (string, error) {
	values := []string{"enabled", "disabled", "test", "mock"}
	return values[m.rand.Intn(len(values))], nil
}

// IntValue returns a random int
func (m *MockConfigService) IntValue(ctx context.Context, key string, params interface{}) (int, error) {
	return m.rand.Intn(1000) + 1, nil
}

// Int32Value returns a random int32
func (m *MockConfigService) Int32Value(ctx context.Context, key string, params interface{}) (int32, error) {
	return int32(m.rand.Intn(1000) + 1), nil
}

// Float64Value returns a random float64
func (m *MockConfigService) Float64Value(ctx context.Context, key string, params interface{}) (float64, error) {
	return m.rand.Float64() * 1000, nil
}

// JsonValue returns random JSON
func (m *MockConfigService) JsonValue(ctx context.Context, key string, params interface{}) (json.RawMessage, error) {
	jsonValues := []string{
		`["STANDARD", "BASIC"]`,
		`{"enabled": true}`,
		`{"config": "test"}`,
		`{"mock": "value"}`,
	}
	return json.RawMessage(jsonValues[m.rand.Intn(len(jsonValues))]), nil
}

// Cluster represents a cluster configuration
type Cluster struct {
	Type string
	ID   string
	Sku  interface{} // Using interface{} to avoid importing product structs
}

// EventsParams represents events parameters
type EventsParams struct {
	AppName       string
	SequenceId    int64
	OrgResourceId string
}

// Connect represents connect configuration
type Connect struct {
	ErrorMappingId int64
}

// ConnectType constant
const ConnectType = "connect"