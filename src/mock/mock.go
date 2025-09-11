package mock

import (
  "context"
  "encoding/json"
  "math/rand"
  "time"
)

// MockConfigService provides mock implementations for config service
type MockConfigService struct {
  rand *rand.Rand
}

// NewMockConfigService creates a new mock config service
func NewMockConfigService() *MockConfigService {
  return &MockConfigService{
    rand: rand.New(rand.NewSource(time.Now().UnixNano())),
  }
}

// BoolValue returns a random boolean value
func (m *MockConfigService) BoolValue(ctx context.Context, key string, params interface{}) (bool, error) {
  return m.rand.Intn(2) == 1, nil
}

// StringValue returns a random string value
func (m *MockConfigService) StringValue(ctx context.Context, key string, params interface{}) (string, error) {
  values := []string{"enabled", "disabled", "test"}
  return values[m.rand.Intn(len(values))], nil
}

// IntValue returns a random int value
func (m *MockConfigService) IntValue(ctx context.Context, key string, params interface{}) (int, error) {
  return m.rand.Intn(1000) + 1, nil
}

// Int32Value returns a random int32 value
func (m *MockConfigService) Int32Value(ctx context.Context, key string, params interface{}) (int32, error) {
  return int32(m.rand.Intn(1000) + 1), nil
}

// Float64Value returns a random float64 value
func (m *MockConfigService) Float64Value(ctx context.Context, key string, params interface{}) (float64, error) {
  return m.rand.Float64() * 1000, nil
}

// JsonValue returns random JSON value
func (m *MockConfigService) JsonValue(ctx context.Context, key string, params interface{}) (json.RawMessage, error) {
  jsonValues := []string{
    `["STANDARD", "BASIC"]`,
    `{"enabled": true}`,
    `{"config": "test"}`,
  }
  return json.RawMessage(jsonValues[m.rand.Intn(len(jsonValues))]), nil
}

// FakeConfigService provides fake implementations
type FakeConfigService struct {
  intValues map[string]int32
  rand      *rand.Rand
}

// NewFakeConfigService creates a new fake config service
func NewFakeConfigService() *FakeConfigService {
  return &FakeConfigService{
    intValues: make(map[string]int32),
    rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
  }
}

// WithInt32Value sets a specific int32 value for a key
func (f *FakeConfigService) WithInt32Value(key string, value int32) *FakeConfigService {
  f.intValues[key] = value
  return f
}

// Int32Value returns the configured or random int32 value
func (f *FakeConfigService) Int32Value(ctx context.Context, key string, params interface{}) (int32, error) {
  if val, exists := f.intValues[key]; exists {
    return val, nil
  }
  return int32(f.rand.Intn(1000) + 1), nil
}
