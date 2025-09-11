package fakeflags

import (
	"context"
	"math/rand"
	"time"
)

// FakeConfigService provides fake implementations for testing
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
