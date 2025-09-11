package idgenmock

import (
  "context"
  "encoding/json"
  "math/rand"
  "time"

  fflags "github.com/confluentinc/cc-fflags"
  "confluent-test/fakeflags"
)

// MockConfigService wraps fake config service for testing
type MockConfigService struct {
  fake *fakeflags.FakeConfigService
  rand *rand.Rand
}

// NewMockConfigService creates a new mock config service with a fake service
func NewMockConfigService(fake *fakeflags.FakeConfigService) *MockConfigService {
  return &MockConfigService{
    fake: fake,
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

// Int32Value returns the configured or random int32 value from fake service
func (m *MockConfigService) Int32Value(ctx context.Context, key string, params interface{}) (int32, error) {
  return m.fake.Int32Value(ctx, key, params)
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