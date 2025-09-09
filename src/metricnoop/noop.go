package metricnoop

import (
	"context"
)

// MeterProvider provides no-op metric implementations
type MeterProvider struct{}

// NewMeterProvider creates a new no-op meter provider
func NewMeterProvider() *MeterProvider {
	return &MeterProvider{}
}

// Meter returns a no-op meter
func (mp *MeterProvider) Meter(name string) *Meter {
	return &Meter{}
}

// Meter provides no-op meter implementations
type Meter struct{}

// Int64Counter returns a no-op counter
func (m *Meter) Int64Counter(name string) (*Counter, error) {
	return &Counter{}, nil
}

// Counter provides no-op counter implementations
type Counter struct{}

// Add does nothing for the no-op counter
func (c *Counter) Add(ctx context.Context, value int64) {
	// No-op implementation
}