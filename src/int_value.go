package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-kit/kit/log"

	fflags "github.com/confluentinc/cc-fflags"
	"github.com/confluentinc/cc-fflags/launchdarkly"
	
	"./flags"
)

func int_value(ld *launchdarkly.LDClient, logger log.Logger) {
	eventsFeatureFlags := fflags.NewEventsConfigService(ld, logger, nil)
	intValue, err := eventsFeatureFlags.IntValue(context.Background(), flags.MyIntFeatureFlag, 
			&fflags.EventsParams{AppName: "test"})
	if err != nil {
		panic(err)
	}
	if intValue == 123 {
		fmt.Println("int value 123")
	} else {
		fmt.Println("int value not 123")
	}
	intValue1, err := eventsFeatureFlags.IntValue(context.Background(), "my-int-feature-flag", 
			&fflags.EventsParams{AppName: "test"})
	if err != nil {
		panic(err)
	}
	if intValue1 == 123 {
		fmt.Println("int value 123")
	} else {
		fmt.Println("int value not 123")
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

	int_value(ld, logger)
} 