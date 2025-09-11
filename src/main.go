package main

import (
  "fmt"
  "os"

  "./launchdarkly"
  "./static"
  "./mock"
  "github.com/go-kit/kit/log"
  fflags "github.com/confluentinc/cc-fflags"
)

func main() {
  fmt.Println("Testing mock dependencies...")
  
  // Test LaunchDarkly mock
  logger := log.NewJSONLogger(os.Stdout)
  ld, err := launchdarkly.New("test-key", 0, logger)
  if err != nil {
    fmt.Printf("LaunchDarkly mock failed: %v\n", err)
    return
  }
  fmt.Printf("LaunchDarkly mock created successfully: %v\n", ld != nil)
  
  // Test static mock
  var ffs *fflags.FFlags
  staticFF := static.NewFeatureFlags(ffs, logger, nil)
  fmt.Printf("Static feature flags mock created successfully: %v\n", staticFF != nil)
  
  // Test mock service
  mockService := mock.NewMockConfigService()
  fmt.Printf("Mock config service created successfully: %v\n", mockService != nil)
  
  fmt.Println("All mock dependencies created successfully!")
}