package static

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"

	fflags "github.com/confluentinc/cc-fflags"
	"github.com/go-kit/kit/log"
)

// FeatureFlags provides static feature flag implementations
type FeatureFlags struct {
	logger log.Logger
	rand   *rand.Rand
}

// NewFeatureFlags creates a new static feature flags instance
func NewFeatureFlags(ffs *fflags.FFlags, logger log.Logger, config interface{}) *FeatureFlags {
	return &FeatureFlags{
		logger: logger,
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// EnableMyCustomFeatureFlag returns a random boolean for custom feature flag
func (ff *FeatureFlags) EnableMyCustomFeatureFlag(ctx context.Context) (bool, error) {
	return ff.rand.Intn(2) == 1, nil
}

// EnableMyCustomIntFeatureFlag returns a random int for custom int feature flag
func (ff *FeatureFlags) EnableMyCustomIntFeatureFlag(ctx context.Context) (int32, error) {
	return int32(ff.rand.Intn(1000) + 1), nil
}

// EnableMyCustomStringFeatureFlag returns a random string for custom string feature flag
func (ff *FeatureFlags) EnableMyCustomStringFeatureFlag(ctx context.Context) (string, error) {
	values := []string{"enabled", "disabled", "test"}
	return values[ff.rand.Intn(len(values))], nil
}

// EnableMyCustomDoubleFeatureFlag returns a random float64 for custom double feature flag
func (ff *FeatureFlags) EnableMyCustomDoubleFeatureFlag(ctx context.Context) (float64, error) {
	return ff.rand.Float64() * 1000, nil
}

// EnableMyCustomJSONFeatureFlag returns random JSON for custom JSON feature flag
func (ff *FeatureFlags) EnableMyCustomJSONFeatureFlag(ctx context.Context) (json.RawMessage, error) {
	jsonValues := []string{
		`{"enabled": true}`,
		`{"config": "test"}`,
		`{"value": 42}`,
	}
	return json.RawMessage(jsonValues[ff.rand.Intn(len(jsonValues))]), nil
}
