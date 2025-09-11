package types

import (
  "context"
  "time"
  "testing"

  fflags "github.com/confluentinc/cc-fflags"
  "confluent-test/metricnoop"
  "confluent-test/idgenmock"
)

// Service interface for the test service
type Service interface {
  SayHello(ctx context.Context, request *SayHelloRequest) (*SayHelloResponse, error)
}

// SayHelloRequest represents a hello request
type SayHelloRequest struct {
  Name string
}

// SayHelloResponse represents a hello response
type SayHelloResponse struct {
  Message  string
  DateTime time.Time
}

// Impl provides service implementation
type Impl struct {
  Clock             Clock
  Params            Params
  ConfigService     fflags.ConfigService
  MeterProvider     *metricnoop.MeterProvider
  DB                interface{}
  WidgetIDGenerator interface{}
}

// Clock interface for time operations
type Clock interface {
  Now() time.Time
}

// RealClock implements Clock interface
type RealClock struct{}

// Now returns current time
func (c *RealClock) Now() time.Time {
  return time.Now()
}

// Params holds service parameters
type Params struct {
  Foo string
}

// MockWidgetIDGenerator creates a mock widget ID generator
func MockWidgetIDGenerator(db interface{}, configService fflags.ConfigService) (interface{}, error) {
  return &MockGenerator{}, nil
}

// MockGenerator provides mock ID generation
type MockGenerator struct{}

// NewClock creates a new real clock
func NewClock() Clock {
  return &RealClock{}
}

// NewParams creates new service parameters
func NewParams() Params {
  return Params{Foo: "test"}
}

// NewDB creates a mock database for testing
func NewDB(t *testing.T) interface{} {
  return &MockDB{}
}

// MockDB provides mock database implementation
type MockDB struct{}

// Suite interface for test suites
type Suite interface {
  T() *testing.T
}

// TestSuite provides basic test suite functionality
type TestSuite struct {
  t *testing.T
}

// T returns the testing.T instance
func (s *TestSuite) T() *testing.T {
  return s.t
}