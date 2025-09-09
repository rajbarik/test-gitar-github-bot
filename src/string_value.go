package main

import (
  "context"
  "fmt"
  "os"
  "time"

  "github.com/go-kit/kit/log"

  fflags "github.com/confluentinc/cc-fflags"
  "./launchdarkly"
  
  "./flags"
)

func string_value(ld *launchdarkly.LDClient, logger log.Logger) {
  eventsFeatureFlags := fflags.NewEventsConfigService(ld, logger, nil)
  key, err := eventsFeatureFlags.StringValue(context.Background(), flags.MyStringFeatureFlag, 
      &fflags.EventsParams{AppName: "test"})
  if err != nil {
    panic(err)
  }
  if key == "enabled" {
    fmt.Println("string value enabled")
  } else {
    fmt.Println("string value disabled")
  }
  key1, err := eventsFeatureFlags.StringValue(context.Background(), "my-str-feature-flag", 
      &fflags.EventsParams{AppName: "test"})
  if err != nil {
    panic(err)
  }
  if key1 == "enabled" {
    fmt.Println("string value enabled")
  } else {
    fmt.Println("string value disabled")
  }
}

func main() {
  logger := log.NewJSONLogger(os.Stdout)
  launchdarklyKey, ok := os.LookupEnv("LAUNCHDARKLY_KEY")
  if !ok {
    fmt.Println("missing required env var LAUNCHDARKLY_KEY")
    os.Exit(1)
  }

  ld, err := launchdarkly.New(launchdarklyKey, 5*time.Second, logger)
  if err != nil {
    panic(err)
  }

  string_value(ld, logger)
} 