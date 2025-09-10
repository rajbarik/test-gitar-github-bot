package launchdarkly

import (
  "math/rand"
  "time"

  log "../go-kit"
)

// LDClient is a mock LaunchDarkly client
type LDClient struct {
  logger log.Logger
  rand   *rand.Rand
}

// New creates a new mock LaunchDarkly client
func New(apiKey string, timeout time.Duration, logger log.Logger) (*LDClient, error) {
  return &LDClient{
    logger: logger,
    rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
  }, nil
}

// Close closes the client connection
func (c *LDClient) Close() error {
  return nil
}

// BoolVariation returns a random boolean value
func (c *LDClient) BoolVariation(key string, user interface{}, defaultValue bool) bool {
  return c.rand.Intn(2) == 1
}

// StringVariation returns a random string value
func (c *LDClient) StringVariation(key string, user interface{}, defaultValue string) string {
  values := []string{"enabled", "disabled", defaultValue}
  return values[c.rand.Intn(len(values))]
}

// IntVariation returns a random int value
func (c *LDClient) IntVariation(key string, user interface{}, defaultValue int) int {
  return c.rand.Intn(1000) + 1
}

// Float64Variation returns a random float64 value
func (c *LDClient) Float64Variation(key string, user interface{}, defaultValue float64) float64 {
  return c.rand.Float64() * 1000
}

// JSONVariation returns a random JSON value
func (c *LDClient) JSONVariation(key string, user interface{}, defaultValue interface{}) interface{} {
  jsonValues := []string{
    `["STANDARD", "BASIC"]`,
    `{"enabled": true}`,
    `{"config": "test"}`,
  }
  return []byte(jsonValues[c.rand.Intn(len(jsonValues))])
}